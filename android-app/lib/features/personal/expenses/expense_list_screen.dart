import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../core/api/expense_api.dart';

class ExpenseListScreen extends ConsumerStatefulWidget {
  const ExpenseListScreen({super.key});

  @override
  ConsumerState<ExpenseListScreen> createState() => _ExpenseListScreenState();
}

class _ExpenseListScreenState extends ConsumerState<ExpenseListScreen> {
  List<dynamic> expenses = [];
  bool isLoading = true;
  String? selectedCategory;

  @override
  void initState() {
    super.initState();
    _loadExpenses();
  }

  Future<void> _loadExpenses() async {
    try {
      final data = await ExpenseApi.getExpenses(filter: selectedCategory);
      setState(() {
        expenses = data;
        isLoading = false;
      });
    } catch (e) {
      setState(() => isLoading = false);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Error: $e')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF020617),
      appBar: AppBar(
        title: const Text('Expenses'),
        backgroundColor: const Color(0xFF1E293B),
        actions: [
          IconButton(
            icon: const Icon(Icons.filter_list),
            onPressed: _showFilterDialog,
          ),
        ],
      ),
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : RefreshIndicator(
              onRefresh: _loadExpenses,
              child: expenses.isEmpty
                  ? const Center(
                      child: Text(
                        'No expenses yet',
                        style: TextStyle(color: Colors.white70),
                      ),
                    )
                  : ListView.builder(
                      padding: const EdgeInsets.all(16),
                      itemCount: expenses.length,
                      itemBuilder: (context, index) {
                        final expense = expenses[index];
                        return _buildExpenseCard(expense);
                      },
                    ),
            ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => _showAddExpenseDialog(),
        backgroundColor: const Color(0xFF3B82F6),
        child: const Icon(Icons.add),
      ),
    );
  }

  Widget _buildExpenseCard(dynamic expense) {
    return Container(
      margin: const EdgeInsets.only(bottom: 12),
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: const Color(0xFF1E293B),
        borderRadius: BorderRadius.circular(12),
      ),
      child: Row(
        children: [
          Container(
            padding: const EdgeInsets.all(12),
            decoration: BoxDecoration(
              color: const Color(0xFF334155),
              borderRadius: BorderRadius.circular(8),
            ),
            child: _getCategoryIcon(expense['category'] ?? ''),
          ),
          const SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  expense['description'] ?? 'Expense',
                  style: const TextStyle(
                    color: Colors.white,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                Text(
                  expense['category'] ?? '',
                  style: const TextStyle(color: Colors.white54),
                ),
                Text(
                  expense['date'] ?? '',
                  style: const TextStyle(color: Colors.white38, fontSize: 12),
                ),
              ],
            ),
          ),
          Text(
            '₹${expense['amount']}',
            style: const TextStyle(
              color: Colors.redAccent,
              fontWeight: FontWeight.bold,
              fontSize: 18,
            ),
          ),
        ],
      ),
    );
  }

  Icon _getCategoryIcon(String category) {
    final icons = {
      'Food & Dining': Icons.restaurant,
      'Groceries': Icons.shopping_cart,
      'Travel & Fuel': Icons.local_gas_station,
      'Shopping': Icons.shopping_bag,
      'Bills & Utilities': Icons.receipt,
      'Health & Wellness': Icons.health_and_safety,
    };
    return Icon(icons[category] ?? Icons.attach_money, color: Colors.white70);
  }

  void _showFilterDialog() {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: const Color(0xFF1E293B),
        title: const Text('Filter by Category', style: TextStyle(color: Colors.white)),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            _filterOption('All Categories', null),
            _filterOption('Food & Dining', 'Food & Dining'),
            _filterOption('Groceries', 'Groceries'),
            _filterOption('Travel & Fuel', 'Travel & Fuel'),
            _filterOption('Shopping', 'Shopping'),
            _filterOption('Bills & Utilities', 'Bills & Utilities'),
          ],
        ),
      ),
    );
  }

  Widget _filterOption(String label, String? category) {
    return ListTile(
      title: Text(label, style: const TextStyle(color: Colors.white)),
      selected: selectedCategory == category,
      onTap: () {
        setState(() => selectedCategory = category);
        _loadExpenses();
        Navigator.pop(context);
      },
    );
  }

  void _showAddExpenseDialog() {
    final amountController = TextEditingController();
    final descController = TextEditingController();
    String selectedCat = 'Food & Dining';

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: const Color(0xFF1E293B),
        title: const Text('Add Expense', style: TextStyle(color: Colors.white)),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: amountController,
              keyboardType: TextInputType.number,
              style: const TextStyle(color: Colors.white),
              decoration: const InputDecoration(
                labelText: 'Amount',
                labelStyle: TextStyle(color: Colors.white70),
              ),
            ),
            TextField(
              controller: descController,
              style: const TextStyle(color: Colors.white),
              decoration: const InputDecoration(
                labelText: 'Description',
                labelStyle: TextStyle(color: Colors.white70),
              ),
            ),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Cancel'),
          ),
          TextButton(
            onPressed: () async {
              try {
                await ExpenseApi.create({
                  'amount': double.parse(amountController.text),
                  'category': selectedCat,
                  'description': descController.text,
                  'date': DateTime.now().toIso8601String().split('T')[0],
                });
                Navigator.pop(context);
                _loadExpenses();
              } catch (e) {
                // Handle error
              }
            },
            child: const Text('Add'),
          ),
        ],
      ),
    );
  }
}

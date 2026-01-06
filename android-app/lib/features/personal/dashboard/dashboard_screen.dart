import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:intl/intl.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../core/api/expense_api.dart';
import '../../../core/api/profile_api.dart';
import '../expenses/expense_list_screen.dart';
import '../reminders/reminder_list_screen.dart';
import '../notes/notes_list_screen.dart';
import '../calendar/calendar_screen.dart';
import '../scanner/scanner_screen.dart';
import '../documents/documents_screen.dart';
import '../../community/hub/community_list_screen.dart';
import '../../chat/chat_screen.dart';
import '../../notification/notification_list_screen.dart';
import './profile_screen.dart';

class DashboardScreen extends ConsumerStatefulWidget {
  const DashboardScreen({super.key});

  @override
  ConsumerState<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends ConsumerState<DashboardScreen> {
  int _selectedIndex = 0;

  final List<Widget> _pages = [
    const PersonalDashboard(),
    const CommunityListScreen(),
    const ChatListScreen(),
    const ProfileScreen(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _pages[_selectedIndex],
      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        backgroundColor: const Color(0xFF1E293B),
        selectedItemColor: const Color(0xFF3B82F6),
        unselectedItemColor: const Color(0xFF94A3B8),
        currentIndex: _selectedIndex,
        onTap: (index) => setState(() => _selectedIndex = index),
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.dashboard),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.people),
            label: 'Community',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.chat),
            label: 'Chat',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.person),
            label: 'Profile',
          ),
        ],
      ),
    );
  }
}

class PersonalDashboard extends StatefulWidget {
  const PersonalDashboard({super.key});

  @override
  State<PersonalDashboard> createState() => _PersonalDashboardState();
}

class _PersonalDashboardState extends State<PersonalDashboard> {
  bool isLoading = true;
  double totalExpenses = 0;
  double monthlyBudget = 50000; // Default, will verify implementation
  List<dynamic> recentTransactions = [];
  String userName = 'User';

  @override
  void initState() {
    super.initState();
    _loadDashboardData();
  }

  Future<void> _loadDashboardData() async {
    try {
      // Load Profile for Name & Budget
      try {
        final profile = await ProfileApi.getProfile();
        if (mounted) {
          setState(() {
            userName = profile['name'] ?? 'User';
            monthlyBudget = (profile['budget'] is int)
                ? (profile['budget'] as int).toDouble()
                : (profile['budget'] as double? ?? 50000.0);
          });
        }
      } catch (e) {
        print('Profile fetch error: $e');
      }

      // Load Expenses for Stats
      final expenses = await ExpenseApi.getExpenses();
      double sum = 0;
      for (var expense in expenses) {
        sum += (expense['amount'] is int) 
            ? (expense['amount'] as int).toDouble() 
            : (expense['amount'] as double? ?? 0.0);
      }

      setState(() {
        totalExpenses = sum;
        recentTransactions = expenses.take(3).toList(); // Top 3
        isLoading = false;
      });
    } catch (e) {
      if (mounted) setState(() => isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    double balance = monthlyBudget - totalExpenses;
    final currencyFormat = NumberFormat.simpleCurrency(name: 'INR', locale: 'en_IN');

    return Scaffold(
      backgroundColor: const Color(0xFF020617),
      appBar: AppBar(
        title: Text('Hello, $userName!'),
        backgroundColor: const Color(0xFF1E293B),
        elevation: 0,
        actions: [
          IconButton(
            icon: const Icon(Icons.notifications_none),
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(builder: (context) => const NotificationListScreen()),
              );
            },
          ),
        ],
      ),
      body: isLoading 
          ? const Center(child: CircularProgressIndicator())
          : SingleChildScrollView(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Balance Card
            Container(
              padding: const EdgeInsets.all(20),
              decoration: BoxDecoration(
                gradient: const LinearGradient(
                  colors: [Color(0xFF3B82F6), Color(0xFF8B5CF6)],
                ),
                borderRadius: BorderRadius.circular(16),
              ),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text(
                    'Total Balance',
                    style: TextStyle(color: Colors.white70),
                  ),
                  const SizedBox(height: 8),
                  Text(
                    currencyFormat.format(balance),
                    style: const TextStyle(
                      color: Colors.white,
                      fontSize: 32,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                  const SizedBox(height: 16),
                  Row(
                    children: [
                      Expanded(
                        child: GestureDetector(
                          onTap: _showEditBudgetDialog,
                          child: Row(
                            children: [
                              _buildBalanceItem('Budget', currencyFormat.format(monthlyBudget), Colors.green),
                              const SizedBox(width: 4),
                              const Icon(Icons.edit, color: Colors.white30, size: 14),
                            ],
                          ),
                        ),
                      ),
                      Expanded(
                        child: _buildBalanceItem('Expenses', currencyFormat.format(totalExpenses), Colors.redAccent),
                      ),
                    ],
                  ),
                ],
              ),
            ),
            const SizedBox(height: 24),
            
            // Quick Actions
            const Text(
              'Quick Actions',
              style: TextStyle(
                color: Colors.white,
                fontSize: 20,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16),
            GridView.count(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              crossAxisCount: 3,
              crossAxisSpacing: 12,
              mainAxisSpacing: 12,
              children: [
                _buildQuickAction(context, Icons.add, 'Add Expense', Colors.orange, () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) => const ExpenseListScreen())).then((_) => _loadDashboardData());
                }),
                _buildQuickAction(context, Icons.camera_alt, 'Scan Bill', Colors.blue, () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) => const ScannerScreen())).then((_) => _loadDashboardData());
                }),
                _buildQuickAction(context, Icons.alarm, 'Reminders', Colors.purple, () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) => const ReminderListScreen()));
                }),
                _buildQuickAction(context, Icons.note, 'Notes', Colors.green, () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) => const NotesListScreen()));
                }),
                _buildQuickAction(context, Icons.folder, 'Documents', Colors.indigo, () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) => const DocumentsScreen()));
                }),
                _buildQuickAction(context, Icons.calendar_today, 'Calendar', Colors.red, () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) => const CalendarScreen()));
                }),
              ],
            ),
            const SizedBox(height: 24),
            
            // Recent Transactions
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                const Text(
                  'Recent Transactions',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 18,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                TextButton(
                  onPressed: () {
                     Navigator.push(context, MaterialPageRoute(builder: (context) => const ExpenseListScreen())).then((_) => _loadDashboardData());
                  },
                  child: const Text('See All'),
                ),
              ],
            ),
            const SizedBox(height: 8),
            if (recentTransactions.isEmpty)
              const Center(child: Padding(
                padding: EdgeInsets.all(16.0),
                child: Text('No recent transactions', style: TextStyle(color: Colors.white38)),
              ))
            else
              ...recentTransactions.map((tx) => _buildTransactionItem(
                _getCategoryIcon(tx['category']),
                tx['description'] ?? 'Expense',
                '-₹${tx['amount']}',
                tx['date'] ?? '',
              )),
          ],
        ),
      ),
    );
  }

  Future<void> _showEditBudgetDialog() async {
    final controller = TextEditingController(text: monthlyBudget.toString());
    await showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: const Color(0xFF1E293B),
        title: const Text('Edit Budget', style: TextStyle(color: Colors.white)),
        content: TextField(
          controller: controller,
          style: const TextStyle(color: Colors.white),
          keyboardType: const TextInputType.numberWithOptions(decimal: true),
          decoration: const InputDecoration(
            labelText: 'Monthly Budget',
            labelStyle: TextStyle(color: Colors.white38),
            enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Colors.white10)),
            focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Color(0xFF3B82F6))),
          ),
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Cancel', style: TextStyle(color: Colors.white38)),
          ),
          ElevatedButton(
            style: ElevatedButton.styleFrom(backgroundColor: const Color(0xFF3B82F6)),
            onPressed: () async {
              final newBudget = double.tryParse(controller.text) ?? monthlyBudget;
              try {
                await ProfileApi.updateProfile({'budget': newBudget});
                if (mounted) {
                  setState(() => monthlyBudget = newBudget);
                  Navigator.pop(context);
                }
              } catch (e) {
                // Handle error
              }
            },
            child: const Text('Save', style: TextStyle(color: Colors.white)),
          ),
        ],
      ),
    );
  }

  Widget _buildBalanceItem(String label, String amount, Color color) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          label,
          style: const TextStyle(color: Colors.white70, fontSize: 12),
        ),
        const SizedBox(height: 4),
        Text(
          amount,
          style: TextStyle(
            color: color,
            fontSize: 18,
            fontWeight: FontWeight.bold,
          ),
        ),
      ],
    );
  }

  Widget _buildQuickAction(BuildContext context, IconData icon, String label, Color color, VoidCallback onTap) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        decoration: BoxDecoration(
          color: const Color(0xFF1E293B),
          borderRadius: BorderRadius.circular(12),
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(icon, color: color, size: 32),
            const SizedBox(height: 8),
            Text(
              label,
              textAlign: TextAlign.center,
              style: const TextStyle(
                color: Colors.white70,
                fontSize: 12,
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildTransactionItem(IconData icon, String title, String amount, String date) {
    return Container(
      margin: const EdgeInsets.only(bottom: 8),
      padding: const EdgeInsets.all(12),
      decoration: BoxDecoration(
        color: const Color(0xFF1E293B),
        borderRadius: BorderRadius.circular(12),
      ),
      child: Row(
        children: [
          Container(
            padding: const EdgeInsets.all(8),
            decoration: BoxDecoration(
              color: const Color(0xFF334155),
              borderRadius: BorderRadius.circular(8),
            ),
            child: Icon(icon, color: Colors.white70),
          ),
          const SizedBox(width: 12),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  title,
                  style: const TextStyle(
                    color: Colors.white,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                Text(
                  date,
                  style: const TextStyle(
                    color: Colors.white54,
                    fontSize: 12,
                  ),
                ),
              ],
            ),
          ),
          Text(
            amount,
            style: const TextStyle(
              color: Colors.red,
              fontWeight: FontWeight.bold,
              fontSize: 16,
            ),
          ),
        ],
      ),
    );
  }
  
  IconData _getCategoryIcon(String? category) {
    final icons = {
      'Food & Dining': Icons.restaurant,
      'Groceries': Icons.shopping_cart,
      'Travel & Fuel': Icons.local_gas_station,
      'Shopping': Icons.shopping_bag,
      'Bills & Utilities': Icons.receipt,
      'Health & Wellness': Icons.health_and_safety,
    };
    return icons[category] ?? Icons.money_off;
  }
}


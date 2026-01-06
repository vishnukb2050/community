import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../core/api/reminder_api.dart';

class ReminderListScreen extends ConsumerStatefulWidget {
  const ReminderListScreen({super.key});

  @override
  ConsumerState<ReminderListScreen> createState() => _ReminderListScreenState();
}

class _ReminderListScreenState extends ConsumerState<ReminderListScreen> {
  List<dynamic> reminders = [];
  bool isLoading = true;
  String filter = 'upcoming';

  @override
  void initState() {
    super.initState();
    _loadReminders();
  }

  Future<void> _loadReminders() async {
    try {
      final data = await ReminderApi.getReminders(filter: filter);
      setState(() {
        reminders = data['reminders'] ?? [];
        isLoading = false;
      });
    } catch (e) {
      setState(() => isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF020617),
      appBar: AppBar(
        title: const Text('Bill Reminders'),
        backgroundColor: const Color(0xFF1E293B),
      ),
      body: Column(
        children: [
          Container(
            padding: const EdgeInsets.all(16),
            child: Row(
              children: [
                _filterChip('Upcoming', 'upcoming'),
                const SizedBox(width: 8),
                _filterChip('Overdue', 'overdue'),
                const SizedBox(width: 8),
                _filterChip('Paid', 'paid'),
              ],
            ),
          ),
          Expanded(
            child: isLoading
                ? const Center(child: CircularProgressIndicator())
                : reminders.isEmpty
                    ? const Center(
                        child: Text('No reminders', style: TextStyle(color: Colors.white70)),
                      )
                    : ListView.builder(
                        padding: const EdgeInsets.all(16),
                        itemCount: reminders.length,
                        itemBuilder: (context, index) {
                          final reminder = reminders[index];
                          return _buildReminderCard(reminder);
                        },
                      ),
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _showAddReminderDialog,
        backgroundColor: const Color(0xFF8B5CF6),
        child: const Icon(Icons.alarm_add),
      ),
    );
  }

  Widget _filterChip(String label, String value) {
    final isSelected = filter == value;
    return GestureDetector(
      onTap: () {
        setState(() => filter = value);
        _loadReminders();
      },
      child: Container(
        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
        decoration: BoxDecoration(
          color: isSelected ? const Color(0xFF8B5CF6) : const Color(0xFF1E293B),
          borderRadius: BorderRadius.circular(20),
        ),
        child: Text(label, style: const TextStyle(color: Colors.white)),
      ),
    );
  }

  Widget _buildReminderCard(dynamic reminder) {
    final status = reminder['status'] ?? 'pending';
    return Container(
      margin: const EdgeInsets.only(bottom: 12),
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: const Color(0xFF1E293B),
        borderRadius: BorderRadius.circular(12),
        border: Border.all(
          color: status == 'overdue' ? Colors.redAccent : Colors.transparent,
          width: 2,
        ),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Expanded(
                child: Text(
                  reminder['title'] ?? '',
                  style: const TextStyle(
                    color: Colors.white,
                    fontWeight: FontWeight.bold,
                    fontSize: 16,
                  ),
                ),
              ),
              Text(
                '₹${reminder['amount']}',
                style: const TextStyle(
                  color: Color(0xFF3B82F6),
                  fontWeight: FontWeight.bold,
                  fontSize: 18,
                ),
              ),
            ],
          ),
          const SizedBox(height: 8),
          Text(
            'Due: ${reminder['due_date']}',
            style: TextStyle(
              color: status == 'overdue' ? Colors.redAccent : Colors.white70,
            ),
          ),
          if (reminder['is_recurring'] == true) ...[
            const SizedBox(height: 4),
            Row(
              children: [
                const Icon(Icons.repeat, size: 14, color: Colors.white38),
                const SizedBox(width: 4),
                Text(
                  reminder['recurrence_type'] ?? 'Recurring',
                  style: const TextStyle(color: Colors.white38, fontSize: 12),
                ),
              ],
            ),
          ],
          const SizedBox(height: 12),
          Row(
            children: [
              if (status != 'paid') ...[
                ElevatedButton(
                  onPressed: () => _markAsPaid(reminder['id']),
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.green,
                    padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
                  ),
                  child: const Text('Mark Paid'),
                ),
                const SizedBox(width: 8),
                OutlinedButton(
                  onPressed: () => _snoozeReminder(reminder['id']),
                  style: OutlinedButton.styleFrom(
                    padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
                  ),
                  child: const Text('Snooze'),
                ),
              ],
            ],
          ),
        ],
      ),
    );
  }

  Future<void> _markAsPaid(String id) async {
    try {
      await ReminderApi.markPaid(id);
      _loadReminders();
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Marked as paid')),
        );
      }
    } catch (e) {
      // Handle error
    }
  }

  Future<void> _snoozeReminder(String id) async {
    try {
      await ReminderApi.snoozeReminder(id, 7); // Snooze for 7 days
      _loadReminders();
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Snoozed for 7 days')),
        );
      }
    } catch (e) {
      // Handle error
    }
  }

  void _showAddReminderDialog() {
    final titleController = TextEditingController();
    final amountController = TextEditingController();
    bool isRecurring = false;

    showDialog(
      context: context,
      builder: (context) => StatefulBuilder(
        builder: (context, setState) => AlertDialog(
          backgroundColor: const Color(0xFF1E293B),
          title: const Text('Add Reminder', style: TextStyle(color: Colors.white)),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              TextField(
                controller: titleController,
                style: const TextStyle(color: Colors.white),
                decoration: const InputDecoration(
                  labelText: 'Title',
                  labelStyle: TextStyle(color: Colors.white70),
                ),
              ),
              TextField(
                controller: amountController,
                keyboardType: TextInputType.number,
                style: const TextStyle(color: Colors.white),
                decoration: const InputDecoration(
                  labelText: 'Amount',
                  labelStyle: TextStyle(color: Colors.white70),
                ),
              ),
              CheckboxListTile(
                title: const Text('Recurring', style: TextStyle(color: Colors.white)),
                value: isRecurring,
                onChanged: (value) => setState(() => isRecurring = value ?? false),
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
                  await ReminderApi.createReminder({
                    'title': titleController.text,
                    'amount': double.parse(amountController.text),
                    'due_date': DateTime.now().add(const Duration(days: 7)).toIso8601String().split('T')[0],
                    'is_recurring': isRecurring,
                    'recurrence_type': isRecurring ? 'monthly' : null,
                  });
                  Navigator.pop(context);
                  this._loadReminders();
                } catch (e) {
                  // Handle error
                }
              },
              child: const Text('Add'),
            ),
          ],
        ),
      ),
    );
  }
}

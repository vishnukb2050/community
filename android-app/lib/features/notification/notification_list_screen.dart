import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../core/api/notification_api.dart';

class NotificationListScreen extends ConsumerStatefulWidget {
  const NotificationListScreen({super.key});

  @override
  ConsumerState<NotificationListScreen> createState() => _NotificationListScreenState();
}

class _NotificationListScreenState extends ConsumerState<NotificationListScreen> {
  List<dynamic> notifications = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    _loadNotifications();
  }

  Future<void> _loadNotifications() async {
    try {
      final data = await NotificationApi.getNotifications();
      setState(() {
        notifications = data['notifications'] ?? [];
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
        title: const Text('Notifications'),
        backgroundColor: const Color(0xFF1E293B),
        actions: [
          TextButton(
            onPressed: () async {
              await NotificationApi.markAllAsRead();
              _loadNotifications();
            },
            child: const Text('Mark all as read'),
          ),
        ],
      ),
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : notifications.isEmpty
              ? const Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Icon(Icons.notifications_none, size: 64, color: Colors.white38),
                      SizedBox(height: 16),
                      Text('No notifications', style: TextStyle(color: Colors.white70)),
                    ],
                  ),
                )
              : ListView.builder(
                  padding: const EdgeInsets.all(16),
                  itemCount: notifications.length,
                  itemBuilder: (context, index) {
                    final notification = notifications[index];
                    return _buildNotificationItem(notification);
                  },
                ),
    );
  }

  Widget _buildNotificationItem(dynamic notification) {
    final isRead = notification['is_read'] ?? false;
    return Container(
      margin: const EdgeInsets.only(bottom: 12),
      decoration: BoxDecoration(
        color: isRead ? const Color(0xFF1E293B) : const Color(0xFF1E293B).withOpacity(0.5),
        borderRadius: BorderRadius.circular(12),
        border: isRead ? null : Border.all(color: Colors.blue.withOpacity(0.5)),
      ),
      child: ListTile(
        leading: CircleAvatar(
          backgroundColor: _getIconColor(notification['type']).withOpacity(0.1),
          child: Icon(_getIcon(notification['type']), color: _getIconColor(notification['type'])),
        ),
        title: Text(
          notification['title'] ?? '',
          style: TextStyle(
            color: Colors.white,
            fontWeight: isRead ? FontWeight.normal : FontWeight.bold,
          ),
        ),
        subtitle: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(notification['message'] ?? '', style: const TextStyle(color: Colors.white70)),
            const SizedBox(height: 4),
            Text(
              notification['created_at'] ?? '',
              style: const TextStyle(color: Colors.white38, fontSize: 10),
            ),
          ],
        ),
        onTap: () async {
          if (!isRead) {
            await NotificationApi.markAsRead(notification['id']);
            _loadNotifications();
          }
        },
      ),
    );
  }

  IconData _getIcon(String? type) {
    switch (type) {
      case 'notice': return Icons.campaign;
      case 'poll': return Icons.how_to_vote;
      case 'meeting': return Icons.groups;
      case 'reminder': return Icons.event_available;
      default: return Icons.notifications;
    }
  }

  Color _getIconColor(String? type) {
    switch (type) {
      case 'notice': return Colors.blue;
      case 'poll': return Colors.green;
      case 'meeting': return Colors.purple;
      case 'reminder': return Colors.orange;
      default: return Colors.blue;
    }
  }
}

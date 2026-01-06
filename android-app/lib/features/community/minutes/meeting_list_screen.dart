import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../core/api/meeting_api.dart';

class MeetingListScreen extends ConsumerStatefulWidget {
  final String communityId;

  const MeetingListScreen({super.key, required this.communityId});

  @override
  ConsumerState<MeetingListScreen> createState() => _MeetingListScreenState();
}

class _MeetingListScreenState extends ConsumerState<MeetingListScreen> {
  List<dynamic> minutes = [];
  bool isLoading = true;
  bool isAdmin = true;

  @override
  void initState() {
    super.initState();
    _loadMinutes();
  }

  Future<void> _loadMinutes() async {
    try {
      final data = await MeetingApi.getMinutes(widget.communityId);
      setState(() {
        minutes = data['minutes'] ?? [];
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
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : minutes.isEmpty
              ? const Center(
                  child: Text('No meetings yet', style: TextStyle(color: Colors.white70)),
                )
              : ListView.builder(
                  padding: const EdgeInsets.all(16),
                  itemCount: minutes.length,
                  itemBuilder: (context, index) {
                    final minute = minutes[index];
                    return _buildMinuteCard(minute);
                  },
                ),
      floatingActionButton: isAdmin
          ? FloatingActionButton(
              onPressed: _showAddMinuteDialog,
              backgroundColor: const Color(0xFF3B82F6),
              child: const Icon(Icons.note_add),
            )
          : null,
    );
  }

  Widget _buildMinuteCard(dynamic minute) {
    return Container(
      margin: const EdgeInsets.only(bottom: 16),
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: const Color(0xFF1E293B),
        borderRadius: BorderRadius.circular(12),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            minute['title'] ?? '',
            style: const TextStyle(
              color: Colors.white,
              fontWeight: FontWeight.bold,
              fontSize: 16,
            ),
          ),
          const SizedBox(height: 8),
          Text(
            'Date: ${minute['meeting_date'] ?? ''}',
            style: const TextStyle(color: Color(0xFF3B82F6)),
          ),
          const SizedBox(height: 12),
          Text(
            minute['content'] ?? '',
            style: const TextStyle(color: Colors.white70),
            maxLines: 3,
            overflow: TextOverflow.ellipsis,
          ),
        ],
      ),
    );
  }

  void _showAddMinuteDialog() {
    final titleController = TextEditingController();
    final contentController = TextEditingController();

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: const Color(0xFF1E293B),
        title: const Text('Add Meeting Minute', style: TextStyle(color: Colors.white)),
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
            const SizedBox(height: 12),
            TextField(
              controller: contentController,
              maxLines: 4,
              style: const TextStyle(color: Colors.white),
              decoration: const InputDecoration(
                labelText: 'Content',
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
                await MeetingApi.createMinute({
                  'community_id': widget.communityId,
                  'title': titleController.text,
                  'content': contentController.text,
                  'meeting_date': DateTime.now().toIso8601String().split('T')[0],
                });
                Navigator.pop(context);
                _loadMinutes();
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

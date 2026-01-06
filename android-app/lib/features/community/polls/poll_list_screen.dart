import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../core/api/poll_api.dart';

class PollListScreen extends ConsumerStatefulWidget {
  final String communityId;

  const PollListScreen({super.key, required this.communityId});

  @override
  ConsumerState<PollListScreen> createState() => _PollListScreenState();
}

class _PollListScreenState extends ConsumerState<PollListScreen> {
  List<dynamic> polls = [];
  bool isLoading = true;
  bool isAdmin = true;

  @override
  void initState() {
    super.initState();
    _loadPolls();
  }

  Future<void> _loadPolls() async {
    try {
      final data = await PollApi.getPolls(widget.communityId);
      setState(() {
        polls = data['polls'] ?? [];
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
          : polls.isEmpty
              ? const Center(
                  child: Text('No polls yet', style: TextStyle(color: Colors.white70)),
                )
              : ListView.builder(
                  padding: const EdgeInsets.all(16),
                  itemCount: polls.length,
                  itemBuilder: (context, index) {
                    final poll = polls[index];
                    return _buildPollCard(poll);
                  },
                ),
      floatingActionButton: isAdmin
          ? FloatingActionButton(
              onPressed: _showCreatePollDialog,
              backgroundColor: const Color(0xFF8B5CF6),
              child: const Icon(Icons.poll),
            )
          : null,
    );
  }

  Widget _buildPollCard(dynamic poll) {
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
            poll['question'] ?? '',
            style: const TextStyle(
              color: Colors.white,
              fontWeight: FontWeight.bold,
              fontSize: 16,
            ),
          ),
          const SizedBox(height: 12),
          ElevatedButton(
            onPressed: () => _showPollOptions(poll['id']),
            child: const Text('View Options & Vote'),
          ),
        ],
      ),
    );
  }

  Future<void> _showPollOptions(String pollId) async {
    try {
      final data = await PollApi.getPollOptions(pollId);
      final options = data['options'] ?? [];

      if (!mounted) return;

      showDialog(
        context: context,
        builder: (context) => AlertDialog(
          backgroundColor: const Color(0xFF1E293B),
          title: const Text('Poll Options', style: TextStyle(color: Colors.white)),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: options.map<Widget>((option) {
              return ListTile(
                title: Text(
                  option['option_text'] ?? '',
                  style: const TextStyle(color: Colors.white),
                ),
                trailing: Text(
                  '${option['vote_count'] ?? 0} votes',
                  style: const TextStyle(color: Colors.white70),
                ),
                onTap: () => _vote(pollId, option['id']),
              );
            }).toList(),
          ),
          actions: [
            TextButton(
              onPressed: () => Navigator.pop(context),
              child: const Text('Close'),
            ),
          ],
        ),
      );
    } catch (e) {
      // Handle error
    }
  }

  Future<void> _vote(String pollId, String optionId) async {
    try {
      await PollApi.vote(pollId, optionId);
      Navigator.pop(context);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Vote recorded!')),
        );
      }
      _loadPolls();
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Error: $e')),
        );
      }
    }
  }

  void _showCreatePollDialog() {
    final questionController = TextEditingController();
    final optionControllers = [TextEditingController(), TextEditingController()];

    showDialog(
      context: context,
      builder: (context) => StatefulBuilder(
        builder: (context, setState) => AlertDialog(
          backgroundColor: const Color(0xFF1E293B),
          title: const Text('Create Poll', style: TextStyle(color: Colors.white)),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              TextField(
                controller: questionController,
                style: const TextStyle(color: Colors.white),
                decoration: const InputDecoration(
                  labelText: 'Question',
                  labelStyle: TextStyle(color: Colors.white70),
                ),
              ),
              const SizedBox(height: 12),
              ...optionControllers.asMap().entries.map((entry) {
                return Padding(
                  padding: const EdgeInsets.only(bottom: 8),
                  child: TextField(
                    controller: entry.value,
                    style: const TextStyle(color: Colors.white),
                    decoration: InputDecoration(
                      labelText: 'Option ${entry.key + 1}',
                      labelStyle: const TextStyle(color: Colors.white70),
                    ),
                  ),
                );
              }).toList(),
              TextButton.icon(
                onPressed: () {
                  setState(() => optionControllers.add(TextEditingController()));
                },
                icon: const Icon(Icons.add),
                label: const Text('Add Option'),
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
                  await PollApi.createPoll({
                    'community_id': widget.communityId,
                    'question': questionController.text,
                    'options': optionControllers.map((c) => c.text).toList(),
                  });
                  Navigator.pop(context);
                  this._loadPolls();
                } catch (e) {
                  // Handle error
                }
              },
              child: const Text('Create'),
            ),
          ],
        ),
      ),
    );
  }
}

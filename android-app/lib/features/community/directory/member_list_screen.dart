import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../core/api/community_api.dart';

class MemberListScreen extends ConsumerStatefulWidget {
  final String communityId;

  const MemberListScreen({super.key, required this.communityId});

  @override
  ConsumerState<MemberListScreen> createState() => _MemberListScreenState();
}

class _MemberListScreenState extends ConsumerState<MemberListScreen> {
  List<dynamic> members = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    _loadMembers();
  }

  Future<void> _loadMembers() async {
    try {
      // Assuming there's a getMembers call or we fetch community detail which includes members
      final data = await CommunityApi.getCommunity(widget.communityId);
      setState(() {
        members = data['members'] ?? [];
        isLoading = false;
      });
    } catch (e) {
      // For now, use mock data if API fails to show structure
      setState(() {
        members = [
          {'name': 'Admin User', 'role': 'admin', 'mobile': '+91 9876543210'},
          {'name': 'John Doe', 'role': 'member', 'mobile': '+91 9123456780'},
        ];
        isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF020617),
      appBar: AppBar(
        title: const Text('Members'),
        backgroundColor: const Color(0xFF1E293B),
      ),
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : members.isEmpty
              ? const Center(
                  child: Text(
                    'No members found',
                    style: TextStyle(color: Colors.white70),
                  ),
                )
              : ListView.separated(
                  padding: const EdgeInsets.all(16),
                  itemCount: members.length,
                  separatorBuilder: (context, index) => const Divider(color: Colors.white10),
                  itemBuilder: (context, index) {
                    final member = members[index];
                    return ListTile(
                      leading: CircleAvatar(
                        backgroundColor: member['role'] == 'admin' ? Colors.blue : const Color(0xFF334155),
                        child: Text(
                          member['name']?.substring(0, 1).toUpperCase() ?? 'U',
                          style: const TextStyle(color: Colors.white),
                        ),
                      ),
                      title: Text(
                        member['name'] ?? 'Unknown',
                        style: const TextStyle(color: Colors.white, fontWeight: FontWeight.bold),
                      ),
                      subtitle: Text(
                        member['mobile'] ?? '',
                        style: const TextStyle(color: Colors.white70),
                      ),
                      trailing: Container(
                        padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                        decoration: BoxDecoration(
                          color: member['role'] == 'admin' ? Colors.blue.withOpacity(0.2) : Colors.white10,
                          borderRadius: BorderRadius.circular(4),
                          border: Border.all(
                            color: member['role'] == 'admin' ? Colors.blue.withOpacity(0.5) : Colors.white24,
                          ),
                        ),
                        child: Text(
                          member['role']?.toUpperCase() ?? 'MEMBER',
                          style: TextStyle(
                            color: member['role'] == 'admin' ? Colors.blue : Colors.white70,
                            fontSize: 10,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ),
                    );
                  },
                ),
      floatingActionButton: FloatingActionButton(
        onPressed: _showAddMemberDialog,
        backgroundColor: const Color(0xFF3B82F6),
        child: const Icon(Icons.person_add),
      ),
    );
  }

  void _showAddMemberDialog() {
    final mobileController = TextEditingController();
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        backgroundColor: const Color(0xFF1E293B),
        title: const Text('Add Member', style: TextStyle(color: Colors.white)),
        content: TextField(
          controller: mobileController,
          style: const TextStyle(color: Colors.white),
          decoration: const InputDecoration(
            labelText: 'Mobile Number',
            labelStyle: TextStyle(color: Colors.white70),
            hintText: 'e.g. 9876543210',
            hintStyle: TextStyle(color: Colors.white38),
          ),
          keyboardType: TextInputType.phone,
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('Cancel', style: TextStyle(color: Colors.white38)),
          ),
          ElevatedButton(
            style: ElevatedButton.styleFrom(backgroundColor: const Color(0xFF3B82F6)),
            onPressed: () async {
              try {
                await CommunityApi.addMember(widget.communityId, mobileController.text);
                Navigator.pop(context);
                _loadMembers();
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Member added successfully')),
                );
              } catch (e) {
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Failed to add member: $e')),
                );
              }
            },
            child: const Text('Add', style: TextStyle(color: Colors.white)),
          ),
        ],
      ),
    );
  }
}

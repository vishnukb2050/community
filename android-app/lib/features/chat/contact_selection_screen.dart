
import 'package:flutter/material.dart';
import 'package:flutter_contacts/flutter_contacts.dart';
import 'package:permission_handler/permission_handler.dart';
import '../../../core/api/auth_api.dart';
import '../../../core/api/chat_api.dart';
import 'chat_screen.dart';

class ContactSelectionScreen extends StatefulWidget {
  const ContactSelectionScreen({super.key});

  @override
  State<ContactSelectionScreen> createState() => _ContactSelectionScreenState();
}

class _ContactSelectionScreenState extends State<ContactSelectionScreen> {
  List<dynamic> appUsers = [];
  bool isLoading = true;
  bool permissionDenied = false;

  @override
  void initState() {
    super.initState();
    _fetchContacts();
  }

  Future<void> _fetchContacts() async {
    if (await Permission.contacts.request().isGranted) {
      try {
        final contacts = await FlutterContacts.getContacts(withProperties: true);
        final numbers = contacts
            .map((c) => c.phones.isNotEmpty ? c.phones.first.number.replaceAll(RegExp(r'\D'), '') : '')
            .where((n) => n.isNotEmpty)
            .toList();

        // Sync with backend
        final result = await AuthApi.syncContacts(numbers);
        setState(() {
          appUsers = result['users'] ?? [];
          isLoading = false;
        });
      } catch (e) {
        // Handle error
        setState(() => isLoading = false);
      }
    } else {
      setState(() {
        permissionDenied = true;
        isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF020617),
      appBar: AppBar(
        title: const Text('New Chat'),
        backgroundColor: const Color(0xFF1E293B),
      ),
      body: isLoading
          ? const Center(child: CircularProgressIndicator())
          : permissionDenied
              ? Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Icon(Icons.perm_contact_calendar_outlined, size: 64, color: Colors.white38),
                      const SizedBox(height: 16),
                      const Text(
                        'Permission denied',
                        style: TextStyle(color: Colors.white70),
                      ),
                      TextButton(
                        onPressed: _fetchContacts,
                        child: const Text('Retry'),
                      ),
                    ],
                  ),
                )
              : appUsers.isEmpty
                  ? const Center(
                      child: Text(
                        'No contacts found on this app',
                        style: TextStyle(color: Colors.white70),
                      ),
                    )
                  : ListView.builder(
                      itemCount: appUsers.length,
                      itemBuilder: (context, index) {
                        final user = appUsers[index];
                        return ListTile(
                          leading: CircleAvatar(
                            backgroundColor: const Color(0xFF3B82F6),
                            child: Text(
                              user['name']?.substring(0, 1).toUpperCase() ?? 'U',
                              style: const TextStyle(color: Colors.white),
                            ),
                          ),
                          title: Text(
                            user['name'] ?? 'Unknown',
                            style: const TextStyle(color: Colors.white, fontWeight: FontWeight.bold),
                          ),
                          subtitle: Text(
                            user['mobile'] ?? '',
                            style: const TextStyle(color: Colors.white70),
                          ),
                          onTap: () => _startChat(user['id']),
                        );
                      },
                    ),
    );
  }

  Future<void> _startChat(String userId) async {
    try {
      final result = await ChatApi.createDirectChat(userId);
      if (!mounted) return;
      
      Navigator.pushReplacement(
        context,
        MaterialPageRoute(
          builder: (context) => ChatScreen(conversationId: result['conversation_id']),
        ),
      );
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Failed to start chat: $e')),
      );
    }
  }
}

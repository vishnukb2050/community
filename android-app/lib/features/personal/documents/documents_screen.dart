import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../core/api/document_api.dart';

class DocumentsScreen extends ConsumerStatefulWidget {
  const DocumentsScreen({super.key});

  @override
  ConsumerState<DocumentsScreen> createState() => _DocumentsScreenState();
}

class _DocumentsScreenState extends ConsumerState<DocumentsScreen> {
  List<dynamic> documents = [];
  List<String> categories = [];
  bool isLoading = true;

  @override
  void initState() {
    super.initState();
    _loadInitialData();
  }

  Future<void> _loadInitialData() async {
    try {
      final docData = await DocumentApi.getDocuments();
      final catData = await DocumentApi.getCategories();
      setState(() {
        documents = docData['documents'] ?? [];
        categories = (catData['categories'] as List?)?.map((e) => e.toString()).toList() ?? [];
        isLoading = false;
      });
    } catch (e) {
      if (mounted) setState(() => isLoading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color(0xFF020617),
      appBar: AppBar(
        title: const Text('Document Vault'),
        backgroundColor: const Color(0xFF1E293B),
        actions: [
          IconButton(
            icon: const Icon(Icons.upload_file),
            onPressed: () {
              // Upload logic
            },
          ),
        ],
      ),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(16),
            child: TextField(
              style: const TextStyle(color: Colors.white),
              decoration: InputDecoration(
                hintText: 'Search documents...',
                hintStyle: const TextStyle(color: Colors.white38),
                prefixIcon: const Icon(Icons.search, color: Colors.white38),
                filled: true,
                fillColor: const Color(0xFF1E293B),
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(12), borderSide: BorderSide.none),
              ),
            ),
          ),
          if (isLoading)
            const Expanded(child: Center(child: CircularProgressIndicator()))
          else
            Expanded(
              child: GridView.builder(
                padding: const EdgeInsets.all(16),
                gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                  crossAxisCount: 2,
                  mainAxisSpacing: 16,
                  crossAxisSpacing: 16,
                ),
                itemCount: categories.isNotEmpty ? categories.length : 6,
                itemBuilder: (context, index) {
                  final cat = categories.isNotEmpty ? categories[index] : 'Untitled Category';
                  final count = documents.where((d) => d['category'] == cat).length;
                  return _buildDocCard(cat, _getIconForCategory(cat), _getColorForCategory(cat), count);
                },
              ),
            ),
        ],
      ),
    );
  }

  Widget _buildDocCard(String title, IconData icon, Color color, int count) {
    return Container(
      decoration: BoxDecoration(
        color: const Color(0xFF1E293B),
        borderRadius: BorderRadius.circular(16),
        border: Border.all(color: Colors.white10),
      ),
      child: Material(
        color: Colors.transparent,
        child: InkWell(
          borderRadius: BorderRadius.circular(16),
          onTap: () {
            // View document category logic
          },
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Icon(icon, size: 40, color: color),
              const SizedBox(height: 12),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 8.0),
                child: Text(
                  title,
                  textAlign: TextAlign.center,
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  style: const TextStyle(color: Colors.white, fontWeight: FontWeight.bold),
                ),
              ),
              Text('$count files', style: const TextStyle(color: Colors.white38, fontSize: 12)),
            ],
          ),
        ),
      ),
    );
  }

  IconData _getIconForCategory(String cat) {
    switch (cat.toLowerCase()) {
      case 'id proof': return Icons.badge;
      case 'bills & receipts': return Icons.receipt_long;
      case 'contracts': return Icons.assignment;
      case 'insurance': return Icons.health_and_safety;
      case 'medical records': return Icons.medical_services;
      default: return Icons.description;
    }
  }

  Color _getColorForCategory(String cat) {
    switch (cat.toLowerCase()) {
      case 'id proof': return Colors.blue;
      case 'bills & receipts': return Colors.orange;
      case 'contracts': return Colors.green;
      case 'insurance': return Colors.red;
      case 'medical records': return Colors.purple;
      default: return Colors.blueGrey;
    }
  }
}

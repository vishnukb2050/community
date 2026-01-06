import 'package:dio/dio.dart';
import 'api_client.dart';

class DocumentApi {
  static Future<Map<String, dynamic>> getDocuments({String? category}) async {
    try {
      final response = await ApiClient.dio.get(
        '/documents',
        queryParameters: category != null ? {'category': category} : null,
      );
      return response.data;
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Failed to fetch documents';
    }
  }

  static Future<Map<String, dynamic>> uploadDocument({
    required String fileName,
    required String category,
    required String base64File,
    String? description,
  }) async {
    try {
      final response = await ApiClient.dio.post(
        '/documents',
        data: {
          'file_name': fileName,
          'category': category,
          'file_base64': base64File,
          'description': description,
        },
      );
      return response.data;
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Upload failed';
    }
  }

  static Future<Map<String, dynamic>> getCategories() async {
    try {
      final response = await ApiClient.dio.get('/documents/categories');
      return response.data;
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Failed to fetch categories';
    }
  }

  static Future<void> deleteDocument(String id) async {
    try {
      await ApiClient.dio.delete('/documents/$id');
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Failed to delete document';
    }
  }
}

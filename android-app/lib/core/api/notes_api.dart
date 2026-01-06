import 'package:dio/dio.dart';
import 'api_client.dart';

class NotesApi {
  static Future<Map<String, dynamic>> getNotes() async {
    final response = await ApiClient.dio.get('/notes');
    return response.data;
  }

  static Future<Map<String, dynamic>> getNote(String id) async {
    final response = await ApiClient.dio.get('/notes/$id');
    return response.data;
  }

  static Future<Map<String, dynamic>> createNote(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/notes', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> updateNote(String id, Map<String, dynamic> data) async {
    final response = await ApiClient.dio.put('/notes/$id', data: data);
    return response.data;
  }

  static Future<void> deleteNote(String id) async {
    await ApiClient.dio.delete('/notes/$id');
  }

  static Future<Map<String, dynamic>> addChecklistItem(String noteId, String itemText) async {
    final response = await ApiClient.dio.post('/notes/$noteId/checklist', data: {'item_text': itemText});
    return response.data;
  }
}

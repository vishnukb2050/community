import 'package:dio/dio.dart';
import 'api_client.dart';

class ChatApi {
  static Future<Map<String, dynamic>> getConversations() async {
    final response = await ApiClient.dio.get('/conversations');
    return response.data;
  }

  static Future<Map<String, dynamic>> createConversation(String participantId) async {
    final response = await ApiClient.dio.post('/conversations', data: {'participant_id': participantId});
    return response.data;
  }

  static Future<Map<String, dynamic>> getMessages(String conversationId) async {
    final response = await ApiClient.dio.get('/conversations/$conversationId/messages');
    return response.data;
  }

  static Future<Map<String, dynamic>> sendMessage(String conversationId, String content) async {
    final response = await ApiClient.dio.post('/conversations/$conversationId/messages', data: {'content': content});
    return response.data;
  }
  static Future<Map<String, dynamic>> createDirectChat(String userId) async {
    final response = await ApiClient.dio.post('/chat/direct', data: {'target_user_id': userId});
    return response.data;
  }
}

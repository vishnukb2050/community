import 'package:dio/dio.dart';
import 'api_client.dart';

class NotificationApi {
  static Future<Map<String, dynamic>> getNotifications({bool unreadOnly = false}) async {
    final params = unreadOnly ? {'unread': 'true'} : null;
    final response = await ApiClient.dio.get('/notifications', queryParameters: params);
    return response.data;
  }

  static Future<Map<String, dynamic>> markAsRead(String id) async {
    final response = await ApiClient.dio.post('/notifications/$id/read');
    return response.data;
  }

  static Future<Map<String, dynamic>> markAllAsRead() async {
    final response = await ApiClient.dio.post('/notifications/read-all');
    return response.data;
  }

  static Future<void> deleteNotification(String id) async {
    await ApiClient.dio.delete('/notifications/$id');
  }

  static Future<Map<String, dynamic>> registerDevice(String deviceToken, String platform) async {
    final response = await ApiClient.dio.post('/notifications/register-device', data: {
      'device_token': deviceToken,
      'platform': platform,
    });
    return response.data;
  }
}

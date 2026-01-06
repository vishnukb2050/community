import 'package:dio/dio.dart';
import 'api_client.dart';

class MeetingApi {
  static Future<Map<String, dynamic>> getMinutes(String communityId) async {
    final response = await ApiClient.dio.get('/minutes', queryParameters: {'community_id': communityId});
    return response.data;
  }

  static Future<Map<String, dynamic>> createMinute(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/minutes', data: data);
    return response.data;
  }

  static Future<void> deleteMinute(String id) async {
    await ApiClient.dio.delete('/minutes/$id');
  }
}

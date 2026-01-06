import 'package:dio/dio.dart';
import 'api_client.dart';

class NoticeApi {
  static Future<Map<String, dynamic>> getNotices(String communityId) async {
    final response = await ApiClient.dio.get('/notices', queryParameters: {'community_id': communityId});
    return response.data;
  }

  static Future<Map<String, dynamic>> getNotice(String id) async {
    final response = await ApiClient.dio.get('/notices/$id');
    return response.data;
  }

  static Future<Map<String, dynamic>> createNotice(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/notices', data: data);
    return response.data;
  }

  static Future<void> deleteNotice(String id) async {
    await ApiClient.dio.delete('/notices/$id');
  }
}

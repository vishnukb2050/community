import 'package:dio/dio.dart';
import 'api_client.dart';

class CommunityApi {
  static Future<Map<String, dynamic>> getCommunities() async {
    final response = await ApiClient.dio.get('/communities');
    return response.data;
  }

  static Future<Map<String, dynamic>> getCommunity(String id) async {
    final response = await ApiClient.dio.get('/communities/$id');
    return response.data;
  }

  static Future<Map<String, dynamic>> createCommunity(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/communities', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> joinCommunity(String inviteCode) async {
    final response = await ApiClient.dio.post('/communities/join', data: {'invite_code': inviteCode});
    return response.data;
  }

  static Future<void> addMember(String communityId, String mobile) async {
    await ApiClient.dio.post('/communities/add-member', data: {
      'community_id': communityId,
      'mobile': mobile,
    });
  }
}

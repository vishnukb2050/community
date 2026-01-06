import 'package:dio/dio.dart';
import 'api_client.dart';

class PollApi {
  static Future<Map<String, dynamic>> getPolls(String communityId) async {
    final response = await ApiClient.dio.get('/polls', queryParameters: {'community_id': communityId});
    return response.data;
  }

  static Future<Map<String, dynamic>> createPoll(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/polls', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> getPollOptions(String pollId) async {
    final response = await ApiClient.dio.get('/polls/$pollId/options');
    return response.data;
  }

  static Future<Map<String, dynamic>> vote(String pollId, String optionId) async {
    final response = await ApiClient.dio.post('/polls/$pollId/vote', data: {'option_id': optionId});
    return response.data;
  }
}

import 'package:dio/dio.dart';
import 'api_client.dart';

class ProfileApi {
  static Future<Map<String, dynamic>> getProfile() async {
    final response = await ApiClient.dio.get('/profile');
    return response.data;
  }

  static Future<Map<String, dynamic>> updateProfile(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.put('/profile', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> updateSettings(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.put('/profile/settings', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> getStats() async {
    final response = await ApiClient.dio.get('/profile/stats');
    return response.data;
  }
}

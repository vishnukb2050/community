import 'package:dio/dio.dart';
import 'api_client.dart';

class ReminderApi {
  static Future<Map<String, dynamic>> getReminders({String? filter}) async {
    final params = filter != null ? {'filter': filter} : null;
    final response = await ApiClient.dio.get('/reminders', queryParameters: params);
    return response.data;
  }

  static Future<Map<String, dynamic>> createReminder(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/reminders', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> updateReminder(String id, Map<String, dynamic> data) async {
    final response = await ApiClient.dio.put('/reminders/$id', data: data);
    return response.data;
  }

  static Future<void> deleteReminder(String id) async {
    await ApiClient.dio.delete('/reminders/$id');
  }

  static Future<Map<String, dynamic>> markPaid(String id) async {
    final response = await ApiClient.dio.post('/reminders/$id/mark-paid');
    return response.data;
  }

  static Future<Map<String, dynamic>> snoozeReminder(String id, int days) async {
    final response = await ApiClient.dio.post('/reminders/$id/snooze', data: {'days': days});
    return response.data;
  }
}

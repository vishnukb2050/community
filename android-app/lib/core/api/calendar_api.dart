import 'package:dio/dio.dart';
import 'api_client.dart';

class CalendarApi {
  static Future<Map<String, dynamic>> getCalendar() async {
    final response = await ApiClient.dio.get('/calendar');
    return response.data;
  }

  static Future<Map<String, dynamic>> getEventsByDate(String date) async {
    final response = await ApiClient.dio.get('/calendar/$date');
    return response.data;
  }

  static Future<Map<String, dynamic>> createEvent(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/calendar/events', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> updateEvent(String id, Map<String, dynamic> data) async {
    final response = await ApiClient.dio.put('/calendar/events/$id', data: data);
    return response.data;
  }

  static Future<void> deleteEvent(String id) async {
    await ApiClient.dio.delete('/calendar/events/$id');
  }
}

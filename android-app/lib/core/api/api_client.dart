import 'package:dio/dio.dart';
import 'package:hive/hive.dart';

class ApiClient {
  static late Dio _dio;
  static const String baseUrl = 'http://13.233.212.142/api'; // EC2 API Gateway (via Nginx)
  
  static void init() {
    _dio = Dio(BaseOptions(
      baseUrl: baseUrl,
      connectTimeout: const Duration(seconds: 30),
      receiveTimeout: const Duration(seconds: 30),
      headers: {
        'Content-Type': 'application/json',
      },
    ));
    
    // Add interceptor for auth token
    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) async {
          final box = await Hive.openBox('settings');
          final token = box.get('auth_token');
          if (token != null) {
            options.headers['Authorization'] = 'Bearer $token';
          }
          return handler.next(options);
        },
        onError: (error, handler) async {
          if (error.response?.statusCode == 401) {
            // Token expired, redirect to login
            final box = await Hive.openBox('settings');
            await box.delete('auth_token');
          }
          return handler.next(error);
        },
      ),
    );
  }
  
  static Dio get dio => _dio;
  
  // Save auth token
  static Future<void> saveToken(String token) async {
    final box = await Hive.openBox('settings');
    await box.put('auth_token', token);
  }
  
  // Get saved token
  static Future<String?> getToken() async {
    final box = await Hive.openBox('settings');
    return box.get('auth_token');
  }
  
  // Clear token (logout)
  static Future<void> clearToken() async {
    final box = await Hive.openBox('settings');
    await box.delete('auth_token');
  }
}

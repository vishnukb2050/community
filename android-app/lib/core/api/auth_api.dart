import 'api_client.dart';

class AuthApi {
  // Send OTP
  static Future<Map<String, dynamic>> sendOTP(String mobile) async {
    final response = await ApiClient.dio.post('/auth/send-otp', data: {
      'mobile': mobile,
    });
    return response.data;
  }
  
  // Verify OTP
  static Future<Map<String, dynamic>> verifyOTP(String mobile, String otp) async {
    final response = await ApiClient.dio.post('/auth/verify-otp', data: {
      'mobile': mobile,
      'otp': otp,
    });
    
    // Save token
    if (response.data['token'] != null) {
      await ApiClient.saveToken(response.data['token']);
    }
    
    return response.data;
  }
  
  // Logout
  static Future<void> logout() async {
    await ApiClient.dio.post('/auth/logout');
    await ApiClient.clearToken();
  }

  static Future<Map<String, dynamic>> syncContacts(List<String> contacts) async {
    final response = await ApiClient.dio.post('/auth/contacts/sync', data: {'contacts': contacts});
    return response.data;
  }
}

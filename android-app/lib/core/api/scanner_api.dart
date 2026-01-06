import 'package:dio/dio.dart';
import 'api_client.dart';

class ScannerApi {
  static Future<Map<String, dynamic>> scanBill(String base64Image) async {
    try {
      final response = await ApiClient.dio.post(
        '/scanner/scan',
        data: {'image_base64': base64Image},
      );
      return response.data;
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Scan failed';
    }
  }

  static Future<Map<String, dynamic>> getScannedBills() async {
    try {
      final response = await ApiClient.dio.get('/scanner/bills');
      return response.data;
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Failed to fetch bills';
    }
  }

  static Future<Map<String, dynamic>> confirmScan(String scanId) async {
    try {
      final response = await ApiClient.dio.post('/scanner/bills/$scanId/confirm');
      return response.data;
    } on DioException catch (e) {
      throw e.response?.data['error'] ?? 'Failed to confirm scan';
    }
  }
}

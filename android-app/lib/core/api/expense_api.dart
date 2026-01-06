import 'api_client.dart';

class ExpenseApi {
  // Get all expenses
  static Future<List<dynamic>> getExpenses({String? filter}) async {
    final response = await ApiClient.dio.get('/expenses', queryParameters: filter != null ? {'filter': filter} : null);
    return response.data['data'] as List;
  }
  
  // Create expense
  static Future<Map<String, dynamic>> createExpense(Map<String, dynamic> data) async {
    final response = await ApiClient.dio.post('/expenses', data: data);
    return response.data;
  }

  static Future<Map<String, dynamic>> create(Map<String, dynamic> data) => createExpense(data);
  
  // Update expense
  static Future<Map<String, dynamic>> updateExpense(String id, Map<String, dynamic> data) async {
    final response = await ApiClient.dio.put('/expenses/$id', data: data);
    return response.data;
  }
  
  // Delete expense
  static Future<void> deleteExpense(String id) async {
    await ApiClient.dio.delete('/expenses/$id');
  }
  
  // Get expense summary
  static Future<Map<String, dynamic>> getSummary() async {
    final response = await ApiClient.dio.get('/expenses/summary');
    return response.data;
  }
  
  // Get categories
  static Future<List<dynamic>> getCategories() async {
    final response = await ApiClient.dio.get('/expenses/categories');
    return response.data['categories'] as List;
  }
}

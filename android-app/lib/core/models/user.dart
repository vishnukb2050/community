class User {
  final String id;
  final String name;
  final String mobile;
  final String? email;
  final String? profilePic;

  User({
    required this.id,
    required this.name,
    required this.mobile,
    this.email,
    this.profilePic,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'],
      name: json['name'],
      mobile: json['mobile'],
      email: json['email'],
      profilePic: json['profile_pic'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'mobile': mobile,
      'email': email,
      'profile_pic': profilePic,
    };
  }
}

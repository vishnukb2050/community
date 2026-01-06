class Community {
  final String id;
  final String name;
  final String? description;
  final String? inviteCode;
  final List<String>? members;

  Community({
    required this.id,
    required this.name,
    this.description,
    this.inviteCode,
    this.members,
  });

  factory Community.fromJson(Map<String, dynamic> json) {
    return Community(
      id: json['id'],
      name: json['name'],
      description: json['description'],
      inviteCode: json['invite_code'],
      members: (json['members'] as List?)?.map((e) => e.toString()).toList(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'invite_code': inviteCode,
      'members': members,
    };
  }
}

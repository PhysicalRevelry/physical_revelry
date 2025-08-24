import 'package:graphql_flutter/graphql_flutter.dart';
import 'package:shared_preferences/shared_preferences.dart';

class GraphQLService {
  static late GraphQLClient _client;

  static Future<void> initialize() async {
    final HttpLink httpLink = HttpLink(
      'http://localhost:8080/graphql', // Use your backend URL
    );

    final AuthLink authLink = AuthLink(
      getToken: () async {
        final prefs = await SharedPreferences.getInstance();
        return 'Bearer ${prefs.getString('auth_token') ?? ''}';
      },
    );

    final Link link = authLink.concat(httpLink);

    _client = GraphQLClient(
      cache: GraphQLCache(store: InMemoryStore()),
      link: link,
    );
  }

  static GraphQLClient get client => _client;
}

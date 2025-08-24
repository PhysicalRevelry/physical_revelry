import 'package:flutter/cupertino.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import 'services/auth_service.dart';
import 'services/graphql_service.dart';
import 'screens/home_screen.dart';
import 'screens/login_screen.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  // Initialize services
  await AuthService.initialize();
  await GraphQLService.initialize();

  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return GraphQLProvider(
      client: ValueNotifier(GraphQLService.client),
      child: CupertinoApp(
        title: 'My App',
        theme: CupertinoThemeData(),
        home: FutureBuilder<bool>(
          future: AuthService.isLoggedIn(),
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const CupertinoActivityIndicator();
            }

            return snapshot.data == true ? const HomeScreen() : const LoginScreen();
          },
        ),
      ),
    );
  }
}

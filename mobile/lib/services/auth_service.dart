import 'dart:developer';

import 'package:auth0_flutter/auth0_flutter.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AuthService {
  static const String _domain = 'YOUR_AUTH0_DOMAIN';
  static const String _clientId = 'YOUR_AUTH0_CLIENT_ID';

  static late Auth0 _auth0;

  static Future<void> initialize() async {
    _auth0 = Auth0(_domain, _clientId);
  }

  static Future<Credentials?> login() async {
    try {
      final credentials = await _auth0.webAuthentication().login();

      // Store token
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('auth_token', credentials.accessToken);

      return credentials;
    } catch (e) {
      log('Login error: $e');
      return null;
    }
  }

  static Future<void> logout() async {
    try {
      await _auth0.webAuthentication().logout();

      // Clear stored token
      final prefs = await SharedPreferences.getInstance();
      await prefs.remove('auth_token');
    } catch (e) {
      log('Logout error: $e');
    }
  }

  static Future<bool> isLoggedIn() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString('auth_token') != null;
  }
}

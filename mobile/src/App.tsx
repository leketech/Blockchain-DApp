import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import Icon from 'react-native-vector-icons/MaterialIcons';

// Screens
import WelcomeScreen from './screens/WelcomeScreen';
import LoginScreen from './screens/LoginScreen';
import RegisterScreen from './screens/RegisterScreen';
import DashboardScreen from './screens/DashboardScreen';
import WalletsScreen from './screens/WalletsScreen';
import SendScreen from './screens/SendScreen';
import ProfileScreen from './screens/ProfileScreen';

// Types
export interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  phone?: string;
  address?: string;
  role: string;
}

export interface Wallet {
  id: number;
  address: string;
  chain: string;
  balance: number;
  is_active: boolean;
}

// Stack and Tab navigators
const Stack = createStackNavigator();
const Tab = createBottomTabNavigator();

// Main tab navigator
const MainTabs = () => {
  return (
    <Tab.Navigator
      screenOptions={({ route }: { route: any }) => ({
        tabBarIcon: ({ focused, color, size }: { focused: boolean; color: string; size: number }) => {
          let iconName = '';
          
          if (route.name === 'Dashboard') {
            iconName = 'dashboard';
          } else if (route.name === 'Wallets') {
            iconName = 'account-balance-wallet';
          } else if (route.name === 'Send') {
            iconName = 'send';
          } else if (route.name === 'Profile') {
            iconName = 'person';
          }
          
          return <Icon name={iconName} size={size} color={color} />;
        },
        tabBarActiveTintColor: '#135bec',
        tabBarInactiveTintColor: 'gray',
        headerStyle: {
          backgroundColor: '#f6f6f8',
        },
        headerTintColor: '#000',
        headerTitleStyle: {
          fontWeight: 'bold',
        },
      })}
    >
      <Tab.Screen 
        name="Dashboard" 
        component={DashboardScreen} 
        options={{ title: 'Dashboard' }} 
      />
      <Tab.Screen 
        name="Wallets" 
        component={WalletsScreen} 
        options={{ title: 'Wallets' }} 
      />
      <Tab.Screen 
        name="Send" 
        component={SendScreen} 
        options={{ title: 'Send' }} 
      />
      <Tab.Screen 
        name="Profile" 
        component={ProfileScreen} 
        options={{ title: 'Profile' }} 
      />
    </Tab.Navigator>
  );
};

// Main app component
const App: React.FC = () => {
  return (
    <NavigationContainer>
      <Stack.Navigator
        initialRouteName="Welcome"
        screenOptions={{
          headerStyle: {
            backgroundColor: '#f6f6f8',
          },
          headerTintColor: '#000',
          headerTitleStyle: {
            fontWeight: 'bold',
          },
        }}
      >
        <Stack.Screen 
          name="Welcome" 
          component={WelcomeScreen} 
          options={{ headerShown: false }} 
        />
        <Stack.Screen 
          name="Login" 
          component={LoginScreen} 
          options={{ title: 'Login' }} 
        />
        <Stack.Screen 
          name="Register" 
          component={RegisterScreen} 
          options={{ title: 'Create Account' }} 
        />
        <Stack.Screen 
          name="Main" 
          component={MainTabs} 
          options={{ headerShown: false }} 
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default App;
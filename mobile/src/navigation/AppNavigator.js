import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import WelcomeScreen from '../screens/WelcomeScreen';

const Stack = createStackNavigator();

const AppNavigator = () => {
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
          options={{ 
            title: 'dApp',
            headerRight: () => (
              <TouchableOpacity onPress={() => console.log('Help pressed')} style={{ marginRight: 16 }}>
                <Text style={{ fontSize: 24, color: '#000' }}>?</Text>
              </TouchableOpacity>
            ),
          }} 
        />
      </Stack.Navigator>
    </NavigationContainer>
  );
};

export default AppNavigator;
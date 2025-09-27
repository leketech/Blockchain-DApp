import React, { useState } from 'react';
import { View, Text, StyleSheet, TextInput, TouchableOpacity, Alert, ScrollView, Switch } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';

interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  phone?: string;
  address?: string;
  role: string;
}

const ProfileScreen: React.FC = () => {
  const [user, setUser] = useState<User>({
    id: 1,
    email: 'user@example.com',
    first_name: 'John',
    last_name: 'Doe',
    phone: '+1 234 567 8900',
    address: '123 Main St, New York, NY 10001',
    role: 'user'
  });

  const [notificationsEnabled, setNotificationsEnabled] = useState(true);
  const [securityAlertsEnabled, setSecurityAlertsEnabled] = useState(true);
  const [isEditing, setIsEditing] = useState(false);

  const handleSave = () => {
    // Here you would implement the actual save logic
    Alert.alert('Success', 'Profile updated successfully!');
    setIsEditing(false);
  };

  const handleLogout = () => {
    Alert.alert(
      'Logout',
      'Are you sure you want to logout?',
      [
        { text: 'Cancel', style: 'cancel' },
        { text: 'Logout', onPress: () => console.log('User logged out') }
      ]
    );
  };

  return (
    <ScrollView style={styles.container}>
      <Text style={styles.title}>Profile</Text>
      
      <View style={styles.profileHeader}>
        <View style={styles.avatarContainer}>
          <View style={styles.avatar}>
            <Text style={styles.avatarText}>
              {user.first_name.charAt(0)}{user.last_name.charAt(0)}
            </Text>
          </View>
          <TouchableOpacity style={styles.editAvatarButton}>
            <Text style={styles.editAvatarText}>Change Photo</Text>
          </TouchableOpacity>
        </View>
        
        <View style={styles.userInfo}>
          <Text style={styles.userName}>
            {user.first_name} {user.last_name}
          </Text>
          <Text style={styles.userEmail}>{user.email}</Text>
          <Text style={styles.userRole}>{user.role.charAt(0).toUpperCase() + user.role.slice(1)}</Text>
        </View>
      </View>

      <View style={styles.section}>
        <View style={styles.sectionHeader}>
          <Text style={styles.sectionTitle}>Personal Information</Text>
          <TouchableOpacity onPress={() => setIsEditing(!isEditing)}>
            <Text style={styles.editText}>
              {isEditing ? 'Cancel' : 'Edit'}
            </Text>
          </TouchableOpacity>
        </View>
        
        <View style={styles.inputGroup}>
          <Text style={styles.label}>First Name</Text>
          {isEditing ? (
            <TextInput
              style={styles.input}
              value={user.first_name}
              onChangeText={(text) => setUser({...user, first_name: text})}
            />
          ) : (
            <Text style={styles.value}>{user.first_name}</Text>
          )}
        </View>
        
        <View style={styles.inputGroup}>
          <Text style={styles.label}>Last Name</Text>
          {isEditing ? (
            <TextInput
              style={styles.input}
              value={user.last_name}
              onChangeText={(text) => setUser({...user, last_name: text})}
            />
          ) : (
            <Text style={styles.value}>{user.last_name}</Text>
          )}
        </View>
        
        <View style={styles.inputGroup}>
          <Text style={styles.label}>Email</Text>
          {isEditing ? (
            <TextInput
              style={styles.input}
              value={user.email}
              onChangeText={(text) => setUser({...user, email: text})}
              keyboardType="email-address"
            />
          ) : (
            <Text style={styles.value}>{user.email}</Text>
          )}
        </View>
        
        <View style={styles.inputGroup}>
          <Text style={styles.label}>Phone</Text>
          {isEditing ? (
            <TextInput
              style={styles.input}
              value={user.phone}
              onChangeText={(text) => setUser({...user, phone: text})}
              keyboardType="phone-pad"
            />
          ) : (
            <Text style={styles.value}>{user.phone}</Text>
          )}
        </View>
        
        <View style={styles.inputGroup}>
          <Text style={styles.label}>Address</Text>
          {isEditing ? (
            <TextInput
              style={[styles.input, styles.textArea]}
              value={user.address}
              onChangeText={(text) => setUser({...user, address: text})}
              multiline
              numberOfLines={3}
            />
          ) : (
            <Text style={[styles.value, styles.multiLineValue]}>{user.address}</Text>
          )}
        </View>
        
        {isEditing && (
          <TouchableOpacity style={styles.saveButton} onPress={handleSave}>
            <Text style={styles.saveButtonText}>Save Changes</Text>
          </TouchableOpacity>
        )}
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Security</Text>
        
        <TouchableOpacity style={styles.settingItem}>
          <View style={styles.settingInfo}>
            <Icon name="lock" size={24} color="#135bec" />
            <Text style={styles.settingText}>Change Password</Text>
          </View>
          <Icon name="chevron-right" size={24} color="#ccc" />
        </TouchableOpacity>
        
        <TouchableOpacity style={styles.settingItem}>
          <View style={styles.settingInfo}>
            <Icon name="security" size={24} color="#135bec" />
            <Text style={styles.settingText}>Two-Factor Authentication</Text>
          </View>
            <Icon name="chevron-right" size={24} color="#ccc" />
        </TouchableOpacity>
        
        <TouchableOpacity style={styles.settingItem}>
          <View style={styles.settingInfo}>
            <Icon name="history" size={24} color="#135bec" />
            <Text style={styles.settingText}>Login History</Text>
          </View>
          <Icon name="chevron-right" size={24} color="#ccc" />
        </TouchableOpacity>
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Notifications</Text>
        
        <View style={styles.notificationSetting}>
          <View style={styles.settingInfo}>
            <Icon name="notifications" size={24} color="#135bec" />
            <Text style={styles.settingText}>Push Notifications</Text>
          </View>
          <Switch
            trackColor={{ false: "#767577", true: "#135bec" }}
            thumbColor={notificationsEnabled ? "#f4f3f4" : "#f4f3f4"}
            ios_backgroundColor="#3e3e3e"
            onValueChange={setNotificationsEnabled}
            value={notificationsEnabled}
          />
        </View>
        
        <View style={styles.notificationSetting}>
          <View style={styles.settingInfo}>
            <Icon name="warning" size={24} color="#135bec" />
            <Text style={styles.settingText}>Security Alerts</Text>
          </View>
          <Switch
            trackColor={{ false: "#767577", true: "#135bec" }}
            thumbColor={securityAlertsEnabled ? "#f4f3f4" : "#f4f3f4"}
            ios_backgroundColor="#3e3e3e"
            onValueChange={setSecurityAlertsEnabled}
            value={securityAlertsEnabled}
          />
        </View>
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Account</Text>
        
        <TouchableOpacity style={styles.settingItem}>
          <View style={styles.settingInfo}>
            <Icon name="help" size={24} color="#135bec" />
            <Text style={styles.settingText}>Help & Support</Text>
          </View>
          <Icon name="chevron-right" size={24} color="#ccc" />
        </TouchableOpacity>
        
        <TouchableOpacity style={styles.settingItem}>
          <View style={styles.settingInfo}>
            <Icon name="info" size={24} color="#135bec" />
            <Text style={styles.settingText}>About</Text>
          </View>
          <Icon name="chevron-right" size={24} color="#ccc" />
        </TouchableOpacity>
        
        <TouchableOpacity style={[styles.settingItem, styles.logoutItem]} onPress={handleLogout}>
          <View style={styles.settingInfo}>
            <Icon name="logout" size={24} color="#ff3b30" />
            <Text style={[styles.settingText, styles.logoutText]}>Logout</Text>
          </View>
        </TouchableOpacity>
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f6f6f8',
    padding: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#000',
    marginBottom: 20,
  },
  profileHeader: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 20,
    marginBottom: 20,
    alignItems: 'center',
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  avatarContainer: {
    alignItems: 'center',
    marginBottom: 16,
  },
  avatar: {
    width: 80,
    height: 80,
    borderRadius: 40,
    backgroundColor: '#135bec',
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 12,
  },
  avatarText: {
    color: '#fff',
    fontSize: 24,
    fontWeight: 'bold',
  },
  editAvatarButton: {
    padding: 8,
  },
  editAvatarText: {
    color: '#135bec',
    fontSize: 14,
    fontWeight: '600',
  },
  userInfo: {
    alignItems: 'center',
  },
  userName: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#000',
    marginBottom: 4,
  },
  userEmail: {
    fontSize: 16,
    color: '#666',
    marginBottom: 4,
  },
  userRole: {
    fontSize: 14,
    color: '#135bec',
    backgroundColor: '#e3eeff',
    paddingHorizontal: 12,
    paddingVertical: 4,
    borderRadius: 12,
  },
  section: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 20,
    marginBottom: 20,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  sectionHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  sectionTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#000',
  },
  editText: {
    color: '#135bec',
    fontSize: 16,
    fontWeight: '600',
  },
  inputGroup: {
    marginBottom: 16,
  },
  label: {
    fontSize: 14,
    color: '#666',
    marginBottom: 8,
  },
  input: {
    backgroundColor: '#f6f6f8',
    borderRadius: 8,
    padding: 12,
    fontSize: 16,
    color: '#000',
  },
  value: {
    fontSize: 16,
    color: '#000',
    paddingVertical: 12,
  },
  multiLineValue: {
    minHeight: 60,
  },
  textArea: {
    height: 80,
    textAlignVertical: 'top',
  },
  saveButton: {
    backgroundColor: '#135bec',
    borderRadius: 8,
    padding: 16,
    alignItems: 'center',
    marginTop: 8,
  },
  saveButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: '600',
  },
  settingItem: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 16,
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
  },
  notificationSetting: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 16,
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
  },
  settingInfo: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  settingText: {
    fontSize: 16,
    color: '#000',
    marginLeft: 16,
  },
  logoutItem: {
    borderBottomWidth: 0,
  },
  logoutText: {
    color: '#ff3b30',
  },
});

export default ProfileScreen;
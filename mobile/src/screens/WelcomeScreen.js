import React from 'react';
import { StyleSheet, Text, View, TouchableOpacity } from 'react-native';

const WelcomeScreen = ({ navigation }) => {
  const handleCreateWallet = () => {
    console.log('Create New Wallet clicked');
    // Navigate to create wallet screen
    // navigation.navigate('CreateWallet');
  };

  const handleImportWallet = () => {
    console.log('Import Existing Wallet clicked');
    // Navigate to import wallet screen
    // navigation.navigate('ImportWallet');
  };

  return (
    <View style={styles.container}>
      {/* Logo */}
      <View style={styles.logoContainer}>
        <View style={styles.logo}>
          <View style={styles.logoInnerCircle} />
          <View style={styles.logoPlus} />
        </View>
      </View>

      {/* Title and Description */}
      <Text style={styles.title}>Welcome to dApp</Text>
      <Text style={styles.description}>
        Your gateway to the decentralized world. Manage your digital assets securely and efficiently.
      </Text>

      {/* Action Buttons */}
      <View style={styles.buttonContainer}>
        <TouchableOpacity style={styles.primaryButton} onPress={handleCreateWallet}>
          <Text style={styles.primaryButtonText}>Create New Wallet</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.secondaryButton} onPress={handleImportWallet}>
          <Text style={styles.secondaryButtonText}>Import Existing Wallet</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f6f6f8',
    justifyContent: 'center',
    alignItems: 'center',
    paddingHorizontal: 24,
  },
  logoContainer: {
    marginBottom: 48,
  },
  logo: {
    width: 80,
    height: 80,
    borderRadius: 16,
    backgroundColor: '#135bec',
    justifyContent: 'center',
    alignItems: 'center',
  },
  logoInnerCircle: {
    width: 40,
    height: 40,
    borderRadius: 20,
    backgroundColor: '#fff',
  },
  logoPlus: {
    position: 'absolute',
    width: 24,
    height: 4,
    backgroundColor: '#fff',
    alignSelf: 'center',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#000',
    textAlign: 'center',
    marginBottom: 16,
  },
  description: {
    fontSize: 16,
    color: '#666',
    textAlign: 'center',
    lineHeight: 24,
    marginBottom: 48,
  },
  buttonContainer: {
    width: '100%',
    gap: 16,
  },
  primaryButton: {
    backgroundColor: '#135bec',
    borderRadius: 8,
    paddingVertical: 16,
    alignItems: 'center',
  },
  primaryButtonText: {
    color: '#fff',
    fontSize: 16,
    fontWeight: 'bold',
  },
  secondaryButton: {
    backgroundColor: 'rgba(19, 91, 236, 0.2)',
    borderRadius: 8,
    paddingVertical: 16,
    alignItems: 'center',
  },
  secondaryButtonText: {
    color: '#135bec',
    fontSize: 16,
    fontWeight: 'bold',
  },
});

export default WelcomeScreen;
import React, { useState } from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Alert, Platform } from 'react-native';
import { Card, DigitalWallet } from '../types/card';

interface AddToWalletScreenProps {
  cards: Card[];
  onAddToWallet: (cardId: string, walletType: string) => void;
  wallets: DigitalWallet[];
}

const AddToWalletScreen: React.FC<AddToWalletScreenProps> = ({ 
  cards, 
  onAddToWallet,
  wallets
}) => {
  const [selectedCard, setSelectedCard] = useState<string>('');
  const [selectedWallet, setSelectedWallet] = useState<string>('apple_pay');

  const handleAddToWallet = () => {
    if (selectedCard && selectedWallet) {
      onAddToWallet(selectedCard, selectedWallet);
      Alert.alert('Success', 'Card added to wallet successfully!');
    } else {
      Alert.alert('Error', 'Please select a card and wallet type');
    }
  };

  const supportedWallets = [
    { id: 'apple_pay', name: 'Apple Pay' },
    { id: 'google_pay', name: 'Google Pay' },
    { id: 'samsung_pay', name: 'Samsung Pay' }
  ];

  // Filter out unsupported wallets based on platform
  const platformSupportedWallets = supportedWallets.filter(wallet => {
    if (Platform.OS === 'ios') {
      return wallet.id === 'apple_pay';
    } else if (Platform.OS === 'android') {
      return wallet.id === 'google_pay' || wallet.id === 'samsung_pay';
    }
    return false;
  });

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Add to Digital Wallet</Text>
      
      {/* Card Selection */}
      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Select Card</Text>
        {cards.map((card) => (
          <TouchableOpacity
            key={card.id}
            style={[
              styles.cardItem,
              selectedCard === card.id && styles.selectedCardItem
            ]}
            onPress={() => setSelectedCard(card.id)}
          >
            <Text style={styles.cardText}>
              {card.brand.toUpperCase()} ****{card.cardNumber?.slice(-4)}
            </Text>
            <Text style={styles.cardSubtext}>
              Expires {card.expiryMonth}/{card.expiryYear}
            </Text>
          </TouchableOpacity>
        ))}
      </View>

      {/* Wallet Selection */}
      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Digital Wallet</Text>
        {platformSupportedWallets.map((wallet) => (
          <TouchableOpacity
            key={wallet.id}
            style={[
              styles.walletItem,
              selectedWallet === wallet.id && styles.selectedWalletItem
            ]}
            onPress={() => setSelectedWallet(wallet.id)}
          >
            <Text style={styles.walletText}>{wallet.name}</Text>
          </TouchableOpacity>
        ))}
      </View>

      {/* Add Button */}
      <TouchableOpacity
        style={styles.addButton}
        onPress={handleAddToWallet}
        disabled={!selectedCard || !selectedWallet}
      >
        <Text style={styles.addButtonText}>Add to Wallet</Text>
      </TouchableOpacity>

      {/* Wallets List */}
      {wallets.length > 0 && (
        <View style={styles.section}>
          <Text style={styles.sectionTitle}>Your Digital Wallets</Text>
          {wallets.map((wallet) => (
            <View key={wallet.id} style={styles.walletEntry}>
              <Text style={styles.walletEntryText}>
                {wallet.walletType.replace('_', ' ')}
              </Text>
              <Text style={styles.walletEntrySubtext}>
                Status: {wallet.status}
              </Text>
            </View>
          ))}
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
    padding: 16,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 24,
    textAlign: 'center',
  },
  section: {
    marginBottom: 24,
  },
  sectionTitle: {
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 12,
  },
  cardItem: {
    backgroundColor: 'white',
    padding: 16,
    marginBottom: 8,
    borderRadius: 8,
    borderWidth: 1,
    borderColor: '#e0e0e0',
  },
  selectedCardItem: {
    borderColor: '#007AFF',
    backgroundColor: '#e6f2ff',
  },
  cardText: {
    fontSize: 16,
    fontWeight: 'bold',
  },
  cardSubtext: {
    fontSize: 14,
    color: '#666',
    marginTop: 4,
  },
  walletItem: {
    backgroundColor: 'white',
    padding: 16,
    marginBottom: 8,
    borderRadius: 8,
    borderWidth: 1,
    borderColor: '#e0e0e0',
  },
  selectedWalletItem: {
    borderColor: '#007AFF',
    backgroundColor: '#e6f2ff',
  },
  walletText: {
    fontSize: 16,
  },
  addButton: {
    backgroundColor: '#007AFF',
    padding: 16,
    borderRadius: 8,
    alignItems: 'center',
    marginBottom: 24,
  },
  addButtonText: {
    color: 'white',
    fontSize: 18,
    fontWeight: 'bold',
  },
  walletEntry: {
    backgroundColor: 'white',
    padding: 16,
    marginBottom: 8,
    borderRadius: 8,
    borderWidth: 1,
    borderColor: '#e0e0e0',
  },
  walletEntryText: {
    fontSize: 16,
    fontWeight: 'bold',
  },
  walletEntrySubtext: {
    fontSize: 14,
    color: '#666',
    marginTop: 4,
  },
});

export default AddToWalletScreen;
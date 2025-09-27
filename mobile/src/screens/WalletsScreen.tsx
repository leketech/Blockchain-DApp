import React, { useState } from 'react';
import { View, Text, StyleSheet, FlatList, TouchableOpacity, Alert } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';

interface Wallet {
  id: number;
  address: string;
  chain: string;
  balance: number;
  is_active: boolean;
}

const WalletsScreen: React.FC = () => {
  const [wallets, setWallets] = useState<Wallet[]>([
    { id: 1, address: 'bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq', chain: 'Bitcoin', balance: 0.5, is_active: true },
    { id: 2, address: '0x742d35Cc6634C0532925a3b8D4C9db4C4C4C4C4C', chain: 'Ethereum', balance: 2.3, is_active: true },
    { id: 3, address: '47htHJdh95mxzbpkgLWlDFGPaPU9ewPRyJMRLCYQiHdQ', chain: 'Solana', balance: 15.7, is_active: true },
  ]);

  const addWallet = () => {
    Alert.alert(
      'Add Wallet',
      'Select a blockchain to create a new wallet',
      [
        { text: 'Bitcoin', onPress: () => console.log('Add Bitcoin wallet') },
        { text: 'Ethereum', onPress: () => console.log('Add Ethereum wallet') },
        { text: 'Solana', onPress: () => console.log('Add Solana wallet') },
        { text: 'Tron', onPress: () => console.log('Add Tron wallet') },
        { text: 'BNB', onPress: () => console.log('Add BNB wallet') },
        { text: 'Cancel', style: 'cancel' },
      ]
    );
  };

  const renderWallet = ({ item }: { item: Wallet }) => (
    <View style={styles.walletCard}>
      <View style={styles.walletHeader}>
        <View style={styles.chainIcon}>
          <Icon name="account-balance" size={24} color="#135bec" />
        </View>
        <View style={styles.walletInfo}>
          <Text style={styles.chainName}>{item.chain}</Text>
          <Text style={styles.address} numberOfLines={1}>
            {item.address}
          </Text>
        </View>
        <View style={styles.balanceContainer}>
          <Text style={styles.balance}>{item.balance.toFixed(4)}</Text>
          <Text style={styles.currency}>USD</Text>
        </View>
      </View>
      <View style={styles.walletActions}>
        <TouchableOpacity style={styles.actionButton}>
          <Icon name="swap-horiz" size={20} color="#135bec" />
          <Text style={styles.actionText}>Swap</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.actionButton}>
          <Icon name="send" size={20} color="#135bec" />
          <Text style={styles.actionText}>Send</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.actionButton}>
          <Icon name="download" size={20} color="#135bec" />
          <Text style={styles.actionText}>Receive</Text>
        </TouchableOpacity>
      </View>
    </View>
  );

  return (
    <View style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.title}>My Wallets</Text>
        <TouchableOpacity style={styles.addButton} onPress={addWallet}>
          <Icon name="add" size={24} color="#fff" />
        </TouchableOpacity>
      </View>
      
      <FlatList
        data={wallets}
        renderItem={renderWallet}
        keyExtractor={(item: Wallet) => item.id.toString()}
        contentContainerStyle={styles.walletList}
        showsVerticalScrollIndicator={false}
      />
      
      <View style={styles.totalBalanceContainer}>
        <Text style={styles.totalLabel}>Total Balance</Text>
        <Text style={styles.totalAmount}>$1,245.67</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f6f6f8',
    paddingTop: 20,
  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingHorizontal: 20,
    marginBottom: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#000',
  },
  addButton: {
    backgroundColor: '#135bec',
    width: 40,
    height: 40,
    borderRadius: 20,
    justifyContent: 'center',
    alignItems: 'center',
  },
  walletList: {
    paddingHorizontal: 20,
  },
  walletCard: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 16,
    marginBottom: 16,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  walletHeader: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 16,
  },
  chainIcon: {
    backgroundColor: '#e3eeff',
    width: 48,
    height: 48,
    borderRadius: 24,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 12,
  },
  walletInfo: {
    flex: 1,
  },
  chainName: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#000',
    marginBottom: 4,
  },
  address: {
    fontSize: 12,
    color: '#666',
    maxWidth: '100%',
  },
  balanceContainer: {
    alignItems: 'flex-end',
  },
  balance: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#000',
  },
  currency: {
    fontSize: 12,
    color: '#666',
  },
  walletActions: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    borderTopWidth: 1,
    borderTopColor: '#eee',
    paddingTop: 16,
  },
  actionButton: {
    alignItems: 'center',
  },
  actionText: {
    fontSize: 12,
    color: '#135bec',
    marginTop: 4,
  },
  totalBalanceContainer: {
    backgroundColor: '#fff',
    padding: 20,
    marginHorizontal: 20,
    borderRadius: 12,
    alignItems: 'center',
    marginBottom: 20,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3,
  },
  totalLabel: {
    fontSize: 16,
    color: '#666',
    marginBottom: 8,
  },
  totalAmount: {
    fontSize: 28,
    fontWeight: 'bold',
    color: '#000',
  },
});

export default WalletsScreen;
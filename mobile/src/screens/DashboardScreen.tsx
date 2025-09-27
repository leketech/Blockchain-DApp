import React, { useState } from 'react';
import { StyleSheet, Text, View, TouchableOpacity, ScrollView, Alert } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';

// Mock data
const mockWallets = [
  { id: 1, address: '0x1234...5678', chain: 'ethereum', balance: 2.5, is_active: true },
  { id: 2, address: 'bc1q...xyz', chain: 'bitcoin', balance: 0.75, is_active: true },
  { id: 3, address: '4t5w...abc', chain: 'solana', balance: 15.2, is_active: true },
];

const DashboardScreen: React.FC = () => {
  const [wallets] = useState(mockWallets);
  
  const totalBalance = wallets.reduce((sum, wallet) => sum + wallet.balance, 0);
  const activeWallets = wallets.filter(wallet => wallet.is_active).length;

  const handleRefresh = () => {
    Alert.alert('Refresh', 'Refreshing wallet data...');
  };

  return (
    <ScrollView style={styles.container}>
      <View style={styles.header}>
        <Text style={styles.headerTitle}>Dashboard</Text>
        <TouchableOpacity onPress={handleRefresh}>
          <Icon name="refresh" size={24} color="#000" />
        </TouchableOpacity>
      </View>
      
      {/* Balance Card */}
      <View style={styles.balanceCard}>
        <Text style={styles.balanceLabel}>Total Balance</Text>
        <Text style={styles.balanceAmount}>${totalBalance.toFixed(2)}</Text>
        <View style={styles.balanceStats}>
          <View style={styles.statItem}>
            <Text style={styles.statValue}>{activeWallets}</Text>
            <Text style={styles.statLabel}>Active Wallets</Text>
          </View>
          <View style={styles.statItem}>
            <Text style={styles.statValue}>{wallets.length}</Text>
            <Text style={styles.statLabel}>Total Wallets</Text>
          </View>
        </View>
      </View>
      
      {/* Quick Actions */}
      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Quick Actions</Text>
        <View style={styles.actionsContainer}>
          <TouchableOpacity style={styles.actionButton}>
            <Icon name="account-balance-wallet" size={24} color="#135bec" />
            <Text style={styles.actionText}>Create Wallet</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.actionButton}>
            <Icon name="send" size={24} color="#135bec" />
            <Text style={styles.actionText}>Send</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.actionButton}>
            <Icon name="swap-horiz" size={24} color="#135bec" />
            <Text style={styles.actionText}>Swap</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.actionButton}>
            <Icon name="history" size={24} color="#135bec" />
            <Text style={styles.actionText}>History</Text>
          </TouchableOpacity>
        </View>
      </View>
      
      {/* Recent Wallets */}
      <View style={styles.section}>
        <View style={styles.sectionHeader}>
          <Text style={styles.sectionTitle}>Your Wallets</Text>
          <TouchableOpacity>
            <Text style={styles.viewAllText}>View All</Text>
          </TouchableOpacity>
        </View>
        
        {wallets.map((wallet) => (
          <View key={wallet.id} style={styles.walletItem}>
            <View style={styles.walletInfo}>
              <View style={styles.walletIcon}>
                <Icon name="account-balance" size={20} color="#fff" />
              </View>
              <View>
                <Text style={styles.walletName}>{wallet.chain.toUpperCase()}</Text>
                <Text style={styles.walletAddress}>{wallet.address}</Text>
              </View>
            </View>
            <Text style={styles.walletBalance}>${wallet.balance.toFixed(2)}</Text>
          </View>
        ))}
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f6f6f8',
  },
  header: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: 24,
    paddingTop: 60,
  },
  headerTitle: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#000',
  },
  balanceCard: {
    backgroundColor: '#135bec',
    margin: 24,
    borderRadius: 16,
    padding: 24,
    alignItems: 'center',
  },
  balanceLabel: {
    fontSize: 16,
    color: 'rgba(255, 255, 255, 0.8)',
    marginBottom: 8,
  },
  balanceAmount: {
    fontSize: 32,
    fontWeight: 'bold',
    color: '#fff',
    marginBottom: 24,
  },
  balanceStats: {
    flexDirection: 'row',
    justifyContent: 'space-around',
    width: '100%',
  },
  statItem: {
    alignItems: 'center',
  },
  statValue: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#fff',
  },
  statLabel: {
    fontSize: 14,
    color: 'rgba(255, 255, 255, 0.8)',
  },
  section: {
    marginHorizontal: 24,
    marginBottom: 24,
  },
  sectionHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 16,
  },
  sectionTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    color: '#000',
  },
  viewAllText: {
    fontSize: 14,
    color: '#135bec',
    fontWeight: '500',
  },
  actionsContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    flexWrap: 'wrap',
  },
  actionButton: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 16,
    alignItems: 'center',
    width: '48%',
    marginBottom: 16,
    elevation: 2,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
  },
  actionText: {
    marginTop: 8,
    fontSize: 14,
    fontWeight: '500',
    color: '#000',
  },
  walletItem: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 16,
    marginBottom: 12,
    elevation: 1,
    shadowColor: '#000',
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.1,
    shadowRadius: 2,
  },
  walletInfo: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  walletIcon: {
    backgroundColor: '#135bec',
    borderRadius: 20,
    width: 40,
    height: 40,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 12,
  },
  walletName: {
    fontSize: 16,
    fontWeight: 'bold',
    color: '#000',
  },
  walletAddress: {
    fontSize: 14,
    color: '#666',
  },
  walletBalance: {
    fontSize: 16,
    fontWeight: 'bold',
    color: '#000',
  },
});

export default DashboardScreen;
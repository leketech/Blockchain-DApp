import React, { useState } from 'react';
import { View, Text, StyleSheet, TextInput, TouchableOpacity, Alert, ScrollView } from 'react-native';
import Icon from 'react-native-vector-icons/MaterialIcons';

interface Wallet {
  id: number;
  address: string;
  chain: string;
  balance: number;
}

const SendScreen: React.FC = () => {
  const [selectedWallet, setSelectedWallet] = useState<Wallet | null>(null);
  const [recipientAddress, setRecipientAddress] = useState('');
  const [amount, setAmount] = useState('');
  const [note, setNote] = useState('');

  const wallets: Wallet[] = [
    { id: 1, address: 'bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq', chain: 'Bitcoin', balance: 0.5 },
    { id: 2, address: '0x742d35Cc6634C0532925a3b8D4C9db4C4C4C4C4C', chain: 'Ethereum', balance: 2.3 },
    { id: 3, address: '47htHJdh95mxzbpkgLWlDFGPaPU9ewPRyJMRLCYQiHdQ', chain: 'Solana', balance: 15.7 },
  ];

  const handleSend = () => {
    if (!selectedWallet) {
      Alert.alert('Error', 'Please select a wallet to send from');
      return;
    }

    if (!recipientAddress) {
      Alert.alert('Error', 'Please enter recipient address');
      return;
    }

    if (!amount || parseFloat(amount) <= 0) {
      Alert.alert('Error', 'Please enter a valid amount');
      return;
    }

    if (parseFloat(amount) > selectedWallet.balance) {
      Alert.alert('Error', 'Insufficient balance');
      return;
    }

    Alert.alert(
      'Confirm Transaction',
      `Send ${amount} ${selectedWallet.chain} to ${recipientAddress}?`,
      [
        { text: 'Cancel', style: 'cancel' },
        { 
          text: 'Send', 
          onPress: () => {
            // Here you would implement the actual send transaction logic
            Alert.alert('Success', 'Transaction sent successfully!');
            setRecipientAddress('');
            setAmount('');
            setNote('');
          }
        },
      ]
    );
  };

  return (
    <ScrollView style={styles.container}>
      <Text style={styles.title}>Send Crypto</Text>
      
      <View style={styles.section}>
        <Text style={styles.sectionTitle}>From Wallet</Text>
        <View style={styles.walletList}>
          {wallets.map((wallet) => (
            <TouchableOpacity
              key={wallet.id}
              style={[
                styles.walletItem,
                selectedWallet?.id === wallet.id && styles.selectedWallet
              ]}
              onPress={() => setSelectedWallet(wallet)}
            >
              <View style={styles.walletInfo}>
                <Text style={styles.walletChain}>{wallet.chain}</Text>
                <Text style={styles.walletAddress} numberOfLines={1}>
                  {wallet.address}
                </Text>
              </View>
              <View style={styles.walletBalance}>
                <Text style={styles.balanceAmount}>{wallet.balance.toFixed(4)}</Text>
                <Text style={styles.balanceCurrency}>{wallet.chain}</Text>
              </View>
              {selectedWallet?.id === wallet.id && (
                <Icon name="check-circle" size={24} color="#135bec" />
              )}
            </TouchableOpacity>
          ))}
        </View>
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>To Address</Text>
        <TextInput
          style={styles.input}
          placeholder="Recipient address"
          value={recipientAddress}
          onChangeText={setRecipientAddress}
          multiline
        />
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Amount</Text>
        <View style={styles.amountContainer}>
          <TextInput
            style={[styles.input, styles.amountInput]}
            placeholder="0.00"
            value={amount}
            onChangeText={setAmount}
            keyboardType="decimal-pad"
          />
          {selectedWallet && (
            <View style={styles.currencyContainer}>
              <Text style={styles.currency}>{selectedWallet.chain}</Text>
            </View>
          )}
        </View>
        <View style={styles.balanceInfo}>
          <Text style={styles.balanceText}>
            Available: {selectedWallet ? selectedWallet.balance.toFixed(4) : '0.00'} {selectedWallet?.chain || ''}
          </Text>
          <TouchableOpacity 
            onPress={() => selectedWallet && setAmount(selectedWallet.balance.toString())}
            disabled={!selectedWallet}
          >
            <Text style={[styles.maxButton, !selectedWallet && styles.disabledMax]}>MAX</Text>
          </TouchableOpacity>
        </View>
      </View>

      <View style={styles.section}>
        <Text style={styles.sectionTitle}>Note (Optional)</Text>
        <TextInput
          style={[styles.input, styles.noteInput]}
          placeholder="Add a note to this transaction"
          value={note}
          onChangeText={setNote}
          multiline
          numberOfLines={3}
        />
      </View>

      <TouchableOpacity 
        style={[styles.sendButton, !selectedWallet && styles.disabledSend]}
        onPress={handleSend}
        disabled={!selectedWallet}
      >
        <Text style={styles.sendButtonText}>Send</Text>
      </TouchableOpacity>
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
  section: {
    marginBottom: 24,
  },
  sectionTitle: {
    fontSize: 16,
    fontWeight: '600',
    color: '#000',
    marginBottom: 12,
  },
  walletList: {
    backgroundColor: '#fff',
    borderRadius: 12,
    overflow: 'hidden',
  },
  walletItem: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 16,
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
  },
  selectedWallet: {
    backgroundColor: '#e3eeff',
  },
  walletInfo: {
    flex: 1,
  },
  walletChain: {
    fontSize: 16,
    fontWeight: '600',
    color: '#000',
    marginBottom: 4,
  },
  walletAddress: {
    fontSize: 12,
    color: '#666',
    maxWidth: '100%',
  },
  walletBalance: {
    alignItems: 'flex-end',
    marginRight: 12,
  },
  balanceAmount: {
    fontSize: 16,
    fontWeight: '600',
    color: '#000',
  },
  balanceCurrency: {
    fontSize: 12,
    color: '#666',
  },
  input: {
    backgroundColor: '#fff',
    borderRadius: 12,
    padding: 16,
    fontSize: 16,
    color: '#000',
    borderWidth: 1,
    borderColor: '#ddd',
  },
  amountContainer: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  amountInput: {
    flex: 1,
    marginRight: 12,
  },
  currencyContainer: {
    backgroundColor: '#135bec',
    borderRadius: 8,
    paddingHorizontal: 16,
    paddingVertical: 12,
  },
  currency: {
    color: '#fff',
    fontWeight: '600',
    fontSize: 16,
  },
  balanceInfo: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginTop: 8,
  },
  balanceText: {
    color: '#666',
    fontSize: 14,
  },
  maxButton: {
    color: '#135bec',
    fontWeight: '600',
    fontSize: 14,
  },
  disabledMax: {
    color: '#ccc',
  },
  noteInput: {
    height: 80,
    textAlignVertical: 'top',
  },
  sendButton: {
    backgroundColor: '#135bec',
    borderRadius: 12,
    padding: 16,
    alignItems: 'center',
    marginTop: 20,
  },
  disabledSend: {
    backgroundColor: '#ccc',
  },
  sendButtonText: {
    color: '#fff',
    fontSize: 18,
    fontWeight: '600',
  },
  disabledSendText: {
    color: '#999',
  },
});

export default SendScreen;
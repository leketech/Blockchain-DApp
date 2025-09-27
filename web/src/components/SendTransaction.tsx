import React, { useState } from 'react';
import { Link } from 'react-router-dom';

interface Wallet {
  id: number;
  address: string;
  chain: string;
  balance: number;
  is_active: boolean;
}

interface SendTransactionProps {
  wallets: Wallet[];
}

const SendTransaction: React.FC<SendTransactionProps> = ({ wallets }) => {
  const [selectedWallet, setSelectedWallet] = useState<number>(wallets[0]?.id || 0);
  const [recipientAddress, setRecipientAddress] = useState('');
  const [amount, setAmount] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    
    // In a real implementation, this would call the backend API
    console.log('Sending transaction:', { selectedWallet, recipientAddress, amount });
    
    // Simulate API call
    setTimeout(() => {
      setLoading(false);
      alert('Transaction sent successfully!');
      setRecipientAddress('');
      setAmount('');
    }, 1500);
  };

  const selectedWalletData = wallets.find(wallet => wallet.id === selectedWallet);

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Navigation */}
      <nav className="bg-white shadow">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex">
              <div className="flex-shrink-0 flex items-center">
                <h1 className="text-xl font-bold text-indigo-600">Blockchain DApp</h1>
              </div>
              <div className="hidden sm:ml-6 sm:flex sm:space-x-8">
                <Link to="/dashboard" className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                  Dashboard
                </Link>
                <Link to="/wallets" className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                  Wallets
                </Link>
                <Link to="/transactions" className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                  Transactions
                </Link>
                <Link to="/send" className="border-indigo-500 text-gray-900 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                  Send
                </Link>
              </div>
            </div>
            <div className="hidden sm:ml-6 sm:flex sm:items-center">
              <Link to="/profile" className="p-1 rounded-full text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                <span className="sr-only">View profile</span>
                <svg className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* Main content */}
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* Page header */}
        <div className="px-4 py-6 sm:px-0">
          <h1 className="text-2xl font-bold text-gray-900">Send Transaction</h1>
          <p className="mt-1 text-gray-600">
            Send cryptocurrency to another address.
          </p>
        </div>

        {/* Send form */}
        <div className="mt-6 bg-white shadow sm:rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <form onSubmit={handleSubmit} className="space-y-6">
              <div>
                <label htmlFor="wallet" className="block text-sm font-medium text-gray-700">
                  From Wallet
                </label>
                <select
                  id="wallet"
                  name="wallet"
                  className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
                  value={selectedWallet}
                  onChange={(e) => setSelectedWallet(Number(e.target.value))}
                >
                  {wallets.map((wallet) => (
                    <option key={wallet.id} value={wallet.id}>
                      {wallet.chain.toUpperCase()} - {wallet.address} (Balance: ${wallet.balance.toFixed(2)})
                    </option>
                  ))}
                </select>
              </div>

              <div>
                <label htmlFor="recipient" className="block text-sm font-medium text-gray-700">
                  Recipient Address
                </label>
                <input
                  type="text"
                  name="recipient"
                  id="recipient"
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  placeholder="Enter recipient address"
                  value={recipientAddress}
                  onChange={(e) => setRecipientAddress(e.target.value)}
                  required
                />
              </div>

              <div>
                <label htmlFor="amount" className="block text-sm font-medium text-gray-700">
                  Amount
                </label>
                <div className="mt-1 relative rounded-md shadow-sm">
                  <input
                    type="number"
                    name="amount"
                    id="amount"
                    className="block w-full pr-12 border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    placeholder="0.00"
                    value={amount}
                    onChange={(e) => setAmount(e.target.value)}
                    step="0.000000000000000001"
                    min="0"
                    required
                  />
                  <div className="absolute inset-y-0 right-0 flex items-center">
                    <label htmlFor="currency" className="sr-only">
                      Currency
                    </label>
                    <select
                      id="currency"
                      name="currency"
                      className="h-full py-0 pl-2 pr-7 border-transparent bg-transparent text-gray-500 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
                      disabled
                    >
                      <option>{selectedWalletData?.chain.toUpperCase() || 'COIN'}</option>
                    </select>
                  </div>
                </div>
                {selectedWalletData && (
                  <p className="mt-2 text-sm text-gray-500">
                    Available balance: ${selectedWalletData.balance.toFixed(2)}
                  </p>
                )}
              </div>

              <div>
                <label htmlFor="fee" className="block text-sm font-medium text-gray-700">
                  Network Fee
                </label>
                <div className="mt-1">
                  <input
                    type="text"
                    name="fee"
                    id="fee"
                    className="block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    value="Standard (0.00021 ETH)"
                    readOnly
                  />
                </div>
              </div>

              <div className="flex items-center">
                <input
                  id="confirm"
                  name="confirm"
                  type="checkbox"
                  className="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                  required
                />
                <label htmlFor="confirm" className="ml-2 block text-sm text-gray-900">
                  I confirm that the recipient address is correct and I want to send this transaction.
                </label>
              </div>

              <div>
                <button
                  type="submit"
                  disabled={loading}
                  className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                  {loading ? (
                    <>
                      <svg className="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                      </svg>
                      Sending...
                    </>
                  ) : (
                    'Send Transaction'
                  )}
                </button>
              </div>
            </form>
          </div>
        </div>

        {/* Security notice */}
        <div className="mt-6 bg-yellow-50 border-l-4 border-yellow-400 p-4">
          <div className="flex">
            <div className="flex-shrink-0">
              <svg className="h-5 w-5 text-yellow-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
              </svg>
            </div>
            <div className="ml-3">
              <p className="text-sm text-yellow-700">
                <strong>Security Notice:</strong> Always double-check the recipient address before sending any cryptocurrency. 
                Transactions on the blockchain are irreversible.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SendTransaction;
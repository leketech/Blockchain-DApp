import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import QRCode from 'qrcode.react';

interface Wallet {
  id: number;
  address: string;
  chain: string;
  balance: number;
  is_active: boolean;
}

interface WalletManagementProps {
  wallets: Wallet[];
  onRefresh: () => void;
}

const WalletManagement: React.FC<WalletManagementProps> = ({ wallets, onRefresh }) => {
  const [showQR, setShowQR] = useState<{[key: number]: boolean}>({});
  const [newWalletChain, setNewWalletChain] = useState('bitcoin');

  const toggleQR = (walletId: number) => {
    setShowQR(prev => ({
      ...prev,
      [walletId]: !prev[walletId]
    }));
  };

  const handleCreateWallet = async () => {
    // In a real implementation, this would call the backend API
    alert(`Creating new ${newWalletChain} wallet`);
    // Refresh the wallet list after creation
    onRefresh();
  };

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
                <Link to="/wallets" className="border-indigo-500 text-gray-900 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                  Wallets
                </Link>
                <Link to="/transactions" className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                  Transactions
                </Link>
                <Link to="/send" className="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
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
          <div className="flex justify-between items-center">
            <h1 className="text-2xl font-bold text-gray-900">Wallet Management</h1>
            <button
              onClick={onRefresh}
              className="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              <svg className="-ml-0.5 mr-2 h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              Refresh
            </button>
          </div>
          <p className="mt-1 text-gray-600">
            Manage your cryptocurrency wallets and view balances.
          </p>
        </div>

        {/* Create wallet form */}
        <div className="mt-6 bg-white shadow sm:rounded-lg">
          <div className="px-4 py-5 sm:p-6">
            <h3 className="text-lg leading-6 font-medium text-gray-900">Create New Wallet</h3>
            <div className="mt-5 sm:flex sm:items-center">
              <div className="w-full sm:max-w-xs">
                <label htmlFor="chain" className="block text-sm font-medium text-gray-700">
                  Blockchain
                </label>
                <select
                  id="chain"
                  name="chain"
                  className="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
                  value={newWalletChain}
                  onChange={(e) => setNewWalletChain(e.target.value)}
                >
                  <option value="bitcoin">Bitcoin</option>
                  <option value="ethereum">Ethereum</option>
                  <option value="solana">Solana</option>
                  <option value="tron">Tron</option>
                  <option value="bnb">BNB</option>
                </select>
              </div>
              <button
                onClick={handleCreateWallet}
                className="mt-3 w-full inline-flex items-center justify-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto"
              >
                Create Wallet
              </button>
            </div>
          </div>
        </div>

        {/* Wallets list */}
        <div className="mt-8">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="max-w-3xl mx-auto">
              <div className="bg-white shadow overflow-hidden sm:rounded-md">
                <ul className="divide-y divide-gray-200">
                  {wallets.map((wallet) => (
                    <li key={wallet.id}>
                      <div className="px-4 py-4 flex items-center sm:px-6">
                        <div className="min-w-0 flex-1 sm:flex sm:items-center sm:justify-between">
                          <div className="truncate">
                            <div className="flex text-sm">
                              <p className="font-medium text-indigo-600 truncate">{wallet.chain.toUpperCase()}</p>
                              <p className="ml-1 flex-shrink-0 font-normal text-gray-500">
                                {wallet.is_active ? (
                                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800">
                                    Active
                                  </span>
                                ) : (
                                  <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800">
                                    Inactive
                                  </span>
                                )}
                              </p>
                            </div>
                            <div className="mt-2 flex">
                              <div className="flex items-center text-sm text-gray-500">
                                <span className="truncate">{wallet.address}</span>
                              </div>
                            </div>
                          </div>
                          <div className="mt-4 flex-shrink-0 sm:mt-0 sm:ml-5">
                            <div className="flex items-center text-sm font-medium text-gray-900">
                              ${wallet.balance.toFixed(2)}
                            </div>
                          </div>
                        </div>
                        <div className="ml-5 flex space-x-2">
                          <button
                            onClick={() => toggleQR(wallet.id)}
                            className="inline-flex items-center px-3 py-1 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                          >
                            QR
                          </button>
                          <button className="inline-flex items-center px-3 py-1 border border-gray-300 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                            <svg className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                            </svg>
                          </button>
                        </div>
                      </div>
                      {showQR[wallet.id] && (
                        <div className="px-4 py-4 bg-gray-50 sm:px-6">
                          <div className="flex justify-center">
                            <QRCode value={wallet.address} size={128} />
                          </div>
                          <div className="mt-2 text-center text-sm text-gray-500">
                            Scan this QR code to receive payments
                          </div>
                        </div>
                      )}
                    </li>
                  ))}
                  {wallets.length === 0 && (
                    <li>
                      <div className="px-4 py-8 text-center sm:px-6">
                        <div className="flex justify-center">
                          <svg className="h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
                          </svg>
                        </div>
                        <h3 className="mt-2 text-sm font-medium text-gray-900">No wallets</h3>
                        <p className="mt-1 text-sm text-gray-500">Get started by creating a new wallet.</p>
                      </div>
                    </li>
                  )}
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default WalletManagement;
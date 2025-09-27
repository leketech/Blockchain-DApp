import React, { useState } from 'react';
import { Card, CardForm, DigitalWallet } from '../types/card';

interface CardProvisioningProps {
  cards: Card[];
  onAddCard: (cardForm: CardForm) => void;
  onAddToWallet: (cardId: string, walletType: string) => void;
  wallets: DigitalWallet[];
}

const CardProvisioning: React.FC<CardProvisioningProps> = ({ 
  cards, 
  onAddCard, 
  onAddToWallet,
  wallets
}) => {
  const [showAddCardForm, setShowAddCardForm] = useState(false);
  const [newCard, setNewCard] = useState<CardForm>({
    cardholderId: '',
    cardType: 'virtual',
    currency: 'usd'
  });
  const [selectedCard, setSelectedCard] = useState<string>('');
  const [selectedWallet, setSelectedWallet] = useState<string>('apple_pay');

  const handleAddCard = () => {
    onAddCard(newCard);
    setNewCard({
      cardholderId: '',
      cardType: 'virtual',
      currency: 'usd'
    });
    setShowAddCardForm(false);
  };

  const handleAddToWallet = () => {
    if (selectedCard && selectedWallet) {
      onAddToWallet(selectedCard, selectedWallet);
    }
  };

  const supportedWallets = [
    { id: 'apple_pay', name: 'Apple Pay' },
    { id: 'google_pay', name: 'Google Pay' },
    { id: 'samsung_pay', name: 'Samsung Pay' }
  ];

  return (
    <div className="card-provisioning">
      <h2 className="text-2xl font-bold mb-4">Card Provisioning</h2>
      
      {/* Add Card Section */}
      <div className="mb-8">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-semibold">Your Cards</h3>
          <button 
            onClick={() => setShowAddCardForm(!showAddCardForm)}
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
          >
            {showAddCardForm ? 'Cancel' : 'Add New Card'}
          </button>
        </div>

        {showAddCardForm && (
          <div className="bg-gray-100 p-4 rounded-lg mb-4">
            <h4 className="text-lg font-medium mb-2">Add New Card</h4>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-700">Card Type</label>
                <select
                  value={newCard.cardType}
                  onChange={(e) => setNewCard({...newCard, cardType: e.target.value})}
                  className="mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                >
                  <option value="virtual">Virtual</option>
                  <option value="physical">Physical</option>
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Currency</label>
                <select
                  value={newCard.currency}
                  onChange={(e) => setNewCard({...newCard, currency: e.target.value})}
                  className="mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                >
                  <option value="usd">USD</option>
                  <option value="eur">EUR</option>
                  <option value="gbp">GBP</option>
                </select>
              </div>
            </div>
            <button
              onClick={handleAddCard}
              className="mt-4 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded"
            >
              Create Card
            </button>
          </div>
        )}

        {/* Cards List */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {cards.map((card) => (
            <div key={card.id} className="bg-white shadow rounded-lg p-4">
              <div className="flex justify-between items-start">
                <div>
                  <h4 className="font-medium">{card.brand.toUpperCase()} {card.type} Card</h4>
                  <p className="text-sm text-gray-500">**** **** **** {card.cardNumber?.slice(-4)}</p>
                  <p className="text-sm">Expires {card.expiryMonth}/{card.expiryYear}</p>
                </div>
                <span className={`px-2 py-1 text-xs rounded-full ${
                  card.status === 'active' ? 'bg-green-100 text-green-800' : 
                  card.status === 'inactive' ? 'bg-yellow-100 text-yellow-800' : 
                  'bg-red-100 text-red-800'
                }`}>
                  {card.status}
                </span>
              </div>
              <div className="mt-4 flex space-x-2">
                <button 
                  onClick={() => {
                    setSelectedCard(card.id);
                    handleAddToWallet();
                  }}
                  className="text-sm bg-blue-100 hover:bg-blue-200 text-blue-800 py-1 px-2 rounded"
                >
                  Add to Wallet
                </button>
                <button 
                  onClick={() => navigator.clipboard.writeText(card.id)}
                  className="text-sm bg-gray-100 hover:bg-gray-200 text-gray-800 py-1 px-2 rounded"
                >
                  Copy ID
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Add to Wallet Section */}
      <div>
        <h3 className="text-xl font-semibold mb-4">Add to Digital Wallet</h3>
        <div className="bg-gray-100 p-4 rounded-lg">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700">Select Card</label>
              <select
                value={selectedCard}
                onChange={(e) => setSelectedCard(e.target.value)}
                className="mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              >
                <option value="">Choose a card</option>
                {cards.map((card) => (
                  <option key={card.id} value={card.id}>
                    {card.brand.toUpperCase()} ****{card.cardNumber?.slice(-4)}
                  </option>
                ))}
              </select>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700">Digital Wallet</label>
              <select
                value={selectedWallet}
                onChange={(e) => setSelectedWallet(e.target.value)}
                className="mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              >
                {supportedWallets.map((wallet) => (
                  <option key={wallet.id} value={wallet.id}>
                    {wallet.name}
                  </option>
                ))}
              </select>
            </div>
          </div>
          <button
            onClick={handleAddToWallet}
            disabled={!selectedCard || !selectedWallet}
            className={`mt-4 ${!selectedCard || !selectedWallet ? 'bg-gray-400' : 'bg-green-500 hover:bg-green-700'} text-white font-bold py-2 px-4 rounded`}
          >
            Add to Wallet
          </button>
        </div>
      </div>

      {/* Wallets List */}
      {wallets.length > 0 && (
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Your Digital Wallets</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {wallets.map((wallet) => (
              <div key={wallet.id} className="bg-white shadow rounded-lg p-4">
                <div className="flex justify-between items-start">
                  <div>
                    <h4 className="font-medium capitalize">{wallet.walletType.replace('_', ' ')}</h4>
                    <p className="text-sm text-gray-500">Device: {wallet.deviceId}</p>
                  </div>
                  <span className={`px-2 py-1 text-xs rounded-full ${
                    wallet.status === 'active' ? 'bg-green-100 text-green-800' : 
                    wallet.status === 'inactive' ? 'bg-yellow-100 text-yellow-800' : 
                    'bg-red-100 text-red-800'
                  }`}>
                    {wallet.status}
                  </span>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default CardProvisioning;
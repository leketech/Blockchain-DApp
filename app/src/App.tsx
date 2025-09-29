import React from 'react';
import WelcomePage from './pages/WelcomePage';
import WalletSetupPage from './pages/WalletSetupPage';
import VirtualCardPage from './pages/VirtualCardPage';

function App() {
  // Simple routing based on hash for now
  const getCurrentPage = () => {
    const hash = window.location.hash;
    
    if (hash === '#/wallet-setup') {
      return <WalletSetupPage />;
    }
    
    if (hash === '#/virtual-card') {
      return <VirtualCardPage />;
    }
    
    if (hash === '#/wallet-created') {
      return <div>Wallet Created - Placeholder</div>;
    }
    
    if (hash === '#/wallet-imported') {
      return <div>Wallet Imported - Placeholder</div>;
    }
    
    if (hash === '#/dashboard') {
      return <div>Dashboard - Placeholder</div>;
    }
    
    return <WelcomePage />;
  };

  return (
    <div className="App">
      {getCurrentPage()}
    </div>
  );
}

export default App;
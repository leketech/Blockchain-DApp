import React from 'react';

const WelcomePage: React.FC = () => {
  const handleCreateWallet = () => {
    console.log('Create New Wallet clicked');
    // Implement wallet creation logic here
  };

  const handleImportWallet = () => {
    console.log('Import Existing Wallet clicked');
    // Implement wallet import logic here
  };

  const handleHelp = () => {
    console.log('Help clicked');
    // Implement help functionality here
  };

  return (
    <div className="flex flex-col min-h-screen">
      <header className="flex items-center justify-between p-4">
        <h2 className="flex-1 text-center text-lg font-bold text-black dark:text-white">dApp</h2>
        <button 
          className="text-black dark:text-white"
          onClick={handleHelp}
          aria-label="Help"
        >
          <span className="material-symbols-outlined">help</span>
        </button>
      </header>
      <main className="flex flex-1 flex-col items-center justify-center p-6 text-center">
        <div className="mb-8">
          <svg fill="none" height="80" viewBox="0 0 80 80" width="80" xmlns="http://www.w3.org/2000/svg">
            <rect fill="url(#paint0_linear_1_2)" height="80" rx="16" width="80"></rect>
            <path d="M40 20C29.074 20 20 29.074 20 40C20 50.926 29.074 60 40 60C50.926 60 60 50.926 60 40C60 29.074 50.926 20 40 20ZM40 56C31.163 56 24 48.837 24 40C24 31.163 31.163 24 40 24C48.837 24 56 31.163 56 40C56 48.837 48.837 56 40 56Z" fill="white"></path>
            <path d="M48 38H42V32H38V38H32V42H38V48H42V42H48V38Z" fill="white"></path>
            <defs>
              <linearGradient gradientUnits="userSpaceOnUse" id="paint0_linear_1_2" x1="40" x2="40" y1="0" y2="80">
                <stop stopColor="#135bec"></stop>
                <stop offset="1" stopColor="#0037a1"></stop>
              </linearGradient>
            </defs>
          </svg>
        </div>
        <h1 className="text-3xl font-bold text-black dark:text-white">Welcome to dApp</h1>
        <p className="mt-2 text-base text-gray-600 dark:text-gray-400">
          Your gateway to the decentralized world. Manage your digital assets securely and efficiently.
        </p>
      </main>
      <footer className="p-6">
        <div className="space-y-4">
          <button 
            className="w-full rounded-lg bg-primary py-3 text-base font-bold text-white transition-opacity hover:opacity-90"
            onClick={handleCreateWallet}
          >
            Create New Wallet
          </button>
          <button 
            className="w-full rounded-lg bg-primary/20 dark:bg-primary/30 py-3 text-base font-bold text-primary transition-opacity hover:opacity-90"
            onClick={handleImportWallet}
          >
            Import Existing Wallet
          </button>
        </div>
      </footer>
    </div>
  );
};

export default WelcomePage;
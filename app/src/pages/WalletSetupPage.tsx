import React from 'react';

const WalletSetupPage: React.FC = () => {
  const handleCreateWallet = () => {
    console.log('Create New Wallet clicked');
    // Implement wallet creation logic here
    // For now, we'll just navigate to a placeholder
    window.location.hash = '#/virtual-card';
  };

  const handleImportWallet = () => {
    console.log('Import Existing Wallet clicked');
    // Implement wallet import logic here
    // For now, we'll just navigate to a placeholder
    window.location.hash = '#/virtual-card';
  };

  return (
    <div className="flex flex-col min-h-screen">
      <header className="flex-shrink-0 p-4">
        <h1 className="text-lg font-bold text-center text-gray-900 dark:text-white">Wallet</h1>
      </header>
      <main className="flex-grow flex flex-col items-center justify-center p-6 text-center">
        <h2 className="text-3xl font-bold text-gray-900 dark:text-white mb-2">Create or import a wallet</h2>
        <p className="text-base text-gray-600 dark:text-gray-400 max-w-sm">Create a new wallet to start your journey or import an existing one to manage your assets.</p>
        <div className="mt-12 w-full max-w-sm space-y-4">
          <button 
            className="flex flex-col items-start p-6 bg-primary/10 dark:bg-primary/20 rounded-xl hover:bg-primary/20 dark:hover:bg-primary/30 transition-colors w-full text-left"
            onClick={handleCreateWallet}
          >
            <h3 className="text-lg font-bold text-gray-900 dark:text-white">Create a new wallet</h3>
            <p className="text-sm text-gray-600 dark:text-gray-400 mt-1">This will create a new, secure wallet and recovery phrase.</p>
          </button>
          <button 
            className="flex flex-col items-start p-6 bg-primary/10 dark:bg-primary/20 rounded-xl hover:bg-primary/20 dark:hover:bg-primary/30 transition-colors w-full text-left"
            onClick={handleImportWallet}
          >
            <h3 className="text-lg font-bold text-gray-900 dark:text-white">Import an existing wallet</h3>
            <p className="text-sm text-gray-600 dark:text-gray-400 mt-1">Import your existing wallet using your secret recovery phrase.</p>
          </button>
        </div>
      </main>
      <footer className="flex-shrink-0 border-t border-primary/20 dark:border-primary/30 bg-background-light dark:bg-background-dark/50 backdrop-blur-sm">
        <nav className="flex justify-around items-center p-2">
          <a className="flex flex-col items-center justify-center gap-1 text-primary w-16 h-16 rounded-full" href="#">
            <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
              <path d="M224,115.55V208a16,16,0,0,1-16,16H168a16,16,0,0,1-16-16V168a8,8,0,0,0-8-8H112a8,8,0,0,0-8,8v40a16,16,0,0,1-16,16H48a16,16,0,0,1-16-16V115.55a16,16,0,0,1,5.17-11.78l80-75.48.11-.11a16,16,0,0,1,21.53,0,1.14,1.14,0,0,0,.11.11l80,75.48A16,16,0,0,1,224,115.55Z"></path>
            </svg>
            <span className="text-xs font-medium">Home</span>
          </a>
          <a className="flex flex-col items-center justify-center gap-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary transition-colors w-16 h-16 rounded-full" href="#">
            <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
              <path d="M216,72H56a8,8,0,0,1,0-16H192a8,8,0,0,0,0-16H56A24,24,0,0,0,32,64V192a24,24,0,0,0,24,24H216a16,16,0,0,0,16-16V88A16,16,0,0,0,216,72Zm0,128H56a8,8,0,0,1-8-8V86.63A23.84,23.84,0,0,0,56,88H216Zm-48-60a12,12,0,1,1,12,12A12,12,0,0,1,168,140Z"></path>
            </svg>
            <span className="text-xs font-medium">Wallet</span>
          </a>
          <a className="flex flex-col items-center justify-center gap-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary transition-colors w-16 h-16 rounded-full" href="#">
            <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
              <path d="M224,48V152a16,16,0,0,1-16,16H99.31l10.35,10.34a8,8,0,0,1-11.32,11.32l-24-24a8,8,0,0,1,0-11.32l24-24a8,8,0,0,1,11.32,11.32L99.31,152H208V48H96v8a8,8,0,0,1-16,0V48A16,16,0,0,1,96,32H208A16,16,0,0,1,224,48ZM168,192a8,8,0,0,0-8,8v8H48V104H156.69l-10.35,10.34a8,8,0,0,0,6.24,2.23l22.82,3.08a16.11,16.11,0,0,0,16-7.86h8.72l3.8,7.86a15.91,15.91,0,0,0,11,8.67l8,1.73L183.86,152h-16.71a16.06,16.06,0,0,0-7.73,2l-12.25,6.76a16.62,16.62,0,0,0-3,2.14l-26.91,24.34A15.93,15.93,0,0,0,110,206.9l.36.65A88.11,88.11,0,0,1,168,192Z"></path>
            </svg>
            <span className="text-xs font-medium">Swap</span>
          </a>
          <a className="flex flex-col items-center justify-center gap-1 text-gray-500 dark:text-gray-400 hover:text-primary dark:hover:text-primary transition-colors w-16 h-16 rounded-full" href="#">
            <svg className="w-6 h-6" fill="currentColor" viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg">
              <path d="M128,24A104,104,0,1,0,232,128,104.11,104.11,0,0,0,128,24Zm88,104a87.62,87.62,0,0,1-6.4,32.94l-44.7-27.49a15.92,15.92,0,0,0-6.24-2.23l-22.82-3.08a16.11,16.11,0,0,0-16,7.86h-8.72l-3.8-7.86a15.91,15.91,0,0,0-11-8.67l-8-1.73L96.14,104h16.71a16.06,16.06,0,0,0,7.73-2l12.25-6.76a16.62,16.62,0,0,0,3-2.14l26.91-24.34A15.93,15.93,0,0,0,166,49.1l-.36-.65A88.11,88.11,0,0,1,216,128ZM143.31,41.34,152,56.9,125.09,81.24,112.85,88H96.14a16,16,0,0,0-13.88,8l-8.73,15.23L63.38,84.19,74.32,58.32a87.87,87.87,0,0,1,69-17ZM40,128a87.53,87.53,0,0,1,8.54-37.8l11.34,30.27a16,16,0,0,0,11.62,10l21.43,4.61L96.74,143a16.09,16.09,0,0,0,14.4,9h1.48l-7.23,16.23a16,16,0,0,0,2.86,17.37l.14.14L128,205.94l-1.94,10A88.11,88.11,0,0,1,40,128Zm102.58,86.78,1.13-5.81a16.09,16.09,0,0,0-4-13.9,1.85,1.85,0,0,1-.14-.14L120,174.74,133.7,144l22.82,3.08,45.72,28.12A88.18,88.18,0,0,1,142.58,214.78Z"></path>
            </svg>
            <span className="text-xs font-medium">Browser</span>
          </a>
        </nav>
      </footer>
    </div>
  );
};

export default WalletSetupPage;
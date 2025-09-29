import React from 'react';

const VirtualCardPage: React.FC = () => {
  const handleAddToAppleWallet = () => {
    console.log('Add to Apple Wallet clicked');
    // Implement Apple Wallet integration
  };

  const handleAddToGoogleWallet = () => {
    console.log('Add to Google Wallet clicked');
    // Implement Google Wallet integration
  };

  const handleContinue = () => {
    console.log('Continue clicked');
    // Navigate to the next step
    window.location.hash = '#/dashboard'; // Placeholder - adjust route as needed
  };

  return (
    <div className="flex flex-col min-h-screen">
      <header className="flex items-center justify-between p-4">
        <h2 className="flex-1 text-center text-lg font-bold text-black dark:text-white">dApp</h2>
        <button className="text-black dark:text-white" aria-label="Help">
          <span className="material-symbols-outlined">help</span>
        </button>
      </header>
      <main className="flex flex-1 flex-col items-center justify-center p-6 text-center">
        <div className="mb-8">
          <div className="relative w-72 h-44 rounded-xl bg-gradient-to-br from-blue-500 to-indigo-700 shadow-lg text-white">
            <div className="p-4">
              <div className="flex justify-between items-center">
                <span className="text-lg font-semibold">dApp</span>
                <span className="material-symbols-outlined text-3xl">contactless</span>
              </div>
              <div className="mt-8">
                <p className="text-sm tracking-wider">•••• •••• •••• 1234</p>
                <p className="text-xs">JOHN APPLESEED</p>
              </div>
            </div>
          </div>
        </div>
        <h1 className="text-3xl font-bold text-black dark:text-white">Your Virtual Debit Card</h1>
        <p className="mt-2 text-base text-gray-600 dark:text-gray-400">
          Ready for Apple & Google Wallet. Your transactions are protected with enhanced security.
        </p>
        <div className="mt-6 flex flex-col items-center space-y-4">
          <button 
            className="flex items-center justify-center w-60 rounded-lg bg-black py-2 text-white transition-opacity hover:opacity-90"
            onClick={handleAddToAppleWallet}
          >
            <img 
              alt="Apple Wallet Logo" 
              className="mr-2" 
              src="https://lh3.googleusercontent.com/aida-public/AB6AXuAXVy_Bfo2iYR1uGNhNww7ES8qqzuekydxETxa5P1Z-lpJBDFdX006woVugj3t5Vv2oJjnQPfgmNH4XrTzWeemRq08zhK-CYY6-dnXf1sYu-IsUuSwICUmfhwAOvWbpFYLRfgGeVHl0DZ3fKuwszhH9zCnepglh_Der2s_8c9dEc7bJSqzfUSLNZPyuhTS0xsOOhb3T_s-3OEKAxI1KJOUTBJll2wf6JEoBEMGpoygtzeVoj4q6oQG-PB7z-jSDLWT3mV2kyw28Wb4" 
            />
            Add to Apple Wallet
          </button>
          <button 
            className="flex items-center justify-center w-60 rounded-lg bg-white dark:bg-zinc-800 py-2 text-black dark:text-white transition-opacity hover:opacity-90 border border-gray-300 dark:border-zinc-700"
            onClick={handleAddToGoogleWallet}
          >
            <img 
              alt="Google Wallet Logo" 
              className="mr-2" 
              src="https://lh3.googleusercontent.com/aida-public/AB6AXuDr2DkK0xQ0A0r3Y5Xc47xRYJKlEL4-hia1PTxCZblYlkHk3Tnt6lrfPNdA6EgSnJIspo-9XHeBBcYlio5jjMiKsImq5GmHZe7KgyNgR-29IM4CgMCT39t8VhHaZ1CBuUfu4amGc-d1bMckxBLQiBnAvBBuOyK1xqLcHpThijNds0xxDKzwPVmXEwOUTnQyDvRB7CPjxB279bbss_OaVdzcsRTr-I_cWHJmeJzpitoOjRSRd6pf2ydaV3IX4iPENv7IpXkkQJhE1g0" 
            />
            Add to Google Wallet
          </button>
        </div>
        <div className="mt-8 flex items-center space-x-2 text-sm text-gray-500 dark:text-gray-400">
          <span className="material-symbols-outlined text-lg">verified_user</span>
          <span>Secured by Google Authenticator</span>
        </div>
      </main>
      <footer className="p-6">
        <div className="space-y-4">
          <button 
            className="w-full rounded-lg bg-primary py-3 text-base font-bold text-white transition-opacity hover:opacity-90"
            onClick={handleContinue}
          >
            Continue
          </button>
        </div>
      </footer>
    </div>
  );
};

export default VirtualCardPage;
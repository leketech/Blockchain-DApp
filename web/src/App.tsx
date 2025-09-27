import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

// Components
import Login from './components/Login';
import Register from './components/Register';
import Dashboard from './components/Dashboard';
import WalletManagement from './components/WalletManagement';
import TransactionHistory from './components/TransactionHistory';
import SendTransaction from './components/SendTransaction';
import Profile from './components/Profile';

// Types
interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  phone?: string;
  address?: string;
  role: string;
}

interface Wallet {
  id: number;
  address: string;
  chain: string;
  balance: number;
  is_active: boolean;
}

const App: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [wallets, setWallets] = useState<Wallet[]>([]);
  const [loading, setLoading] = useState(true);

  // Check if user is authenticated on app load
  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      // Verify token and fetch user data
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      fetchUserData();
    } else {
      setLoading(false);
    }
  }, []);

  const fetchUserData = async () => {
    try {
      const response = await axios.get('/api/v1/auth/user');
      setUser(response.data.user);
      fetchWallets();
    } catch (error) {
      console.error('Failed to fetch user data:', error);
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];
    } finally {
      setLoading(false);
    }
  };

  const fetchWallets = async () => {
    try {
      const response = await axios.get('/api/v1/wallets');
      setWallets(response.data);
    } catch (error) {
      console.error('Failed to fetch wallets:', error);
      toast.error('Failed to load wallets');
    }
  };

  const handleLogin = async (email: string, password: string) => {
    try {
      const response = await axios.post('/api/v1/auth/login', { email, password });
      const { user, token } = response.data;
      
      localStorage.setItem('token', token);
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      setUser(user);
      
      toast.success('Login successful');
      return true;
    } catch (error) {
      toast.error('Login failed. Please check your credentials.');
      return false;
    }
  };

  const handleRegister = async (email: string, password: string, firstName: string, lastName: string) => {
    try {
      await axios.post('/api/v1/auth/register', { email, password, first_name: firstName, last_name: lastName });
      toast.success('Registration successful. Please login.');
      return true;
    } catch (error) {
      toast.error('Registration failed. Please try again.');
      return false;
    }
  };

  const handleLogout = async () => {
    try {
      await axios.post('/api/v1/auth/logout');
    } catch (error) {
      console.error('Logout error:', error);
    } finally {
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];
      setUser(null);
      setWallets([]);
      toast.info('You have been logged out');
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <ToastContainer position="top-right" autoClose={3000} />
        <Routes>
          <Route path="/login" element={user ? <Navigate to="/dashboard" /> : <Login onLogin={handleLogin} />} />
          <Route path="/register" element={user ? <Navigate to="/dashboard" /> : <Register onRegister={handleRegister} />} />
          <Route path="/dashboard" element={user ? <Dashboard user={user} wallets={wallets} onLogout={handleLogout} /> : <Navigate to="/login" />} />
          <Route path="/wallets" element={user ? <WalletManagement wallets={wallets} onRefresh={fetchWallets} /> : <Navigate to="/login" />} />
          <Route path="/transactions" element={user ? <TransactionHistory /> : <Navigate to="/login" />} />
          <Route path="/send" element={user ? <SendTransaction wallets={wallets} /> : <Navigate to="/login" />} />
          <Route path="/profile" element={user ? <Profile user={user} /> : <Navigate to="/login" />} />
          <Route path="/" element={<Navigate to={user ? "/dashboard" : "/login"} />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
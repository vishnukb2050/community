import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import Layout from './components/Layout';
import Login from './pages/Login';
import Dashboard from './pages/Dashboard';
import Notices from './pages/Notices';
import Polls from './pages/Polls';
import Meetings from './pages/Meetings';
import Members from './pages/Members';
import MobileDemo from './pages/MobileDemo';

function App() {
    return (
        <AuthProvider>
            <Router>
                <Routes>
                    <Route path="/demo" element={<MobileDemo />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/" element={<Layout />}>
                        <Route index element={<Dashboard />} />
                        <Route path="notices" element={<Notices />} />
                        <Route path="polls" element={<Polls />} />
                        <Route path="meetings" element={<Meetings />} />
                        <Route path="members" element={<Members />} />
                    </Route>
                    <Route path="*" element={<Navigate to="/" replace />} />
                </Routes>
            </Router>
        </AuthProvider>
    );
}

export default App;

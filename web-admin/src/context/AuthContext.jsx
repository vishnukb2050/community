import React, { createContext, useContext, useState, useEffect } from 'react';
import { authService } from '../services/communityApi';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [communityId, setCommunityId] = useState(localStorage.getItem('selected_community') || '');

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            // In production, verify token or fetch profile
            setUser({ token });
        }
        setLoading(false);
    }, []);

    const login = async (mobile, otp) => {
        const response = await authService.verifyOTP(mobile, otp);
        localStorage.setItem('token', response.token);
        setUser({ token: response.token });
        return response;
    };

    const logout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('selected_community');
        setUser(null);
        setCommunityId('');
    };

    const selectCommunity = (id) => {
        setCommunityId(id);
        localStorage.setItem('selected_community', id);
    };

    return (
        <AuthContext.Provider value={{ user, loading, login, logout, communityId, selectCommunity }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);

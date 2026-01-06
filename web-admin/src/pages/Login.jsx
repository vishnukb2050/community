import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { authService } from '../services/communityApi';

const Login = () => {
    const [mobile, setMobile] = useState('');
    const [otp, setOtp] = useState('');
    const [step, setStep] = useState(1);
    const [error, setError] = useState('');
    const [receivedOtp, setReceivedOtp] = useState('');
    const { login } = useAuth();
    const navigate = useNavigate();

    const handleSendOTP = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        try {
            const data = await authService.sendOTP(mobile);
            if (data.otp) {
                setReceivedOtp(data.otp);
            }
            setStep(2);
        } catch (err) {
            setError('Failed to send OTP. Please check your number.');
        } finally {
            setLoading(false);
        }
    };

    const handleVerifyOTP = async (e) => {
        e.preventDefault();
        setLoading(true);
        setError('');
        try {
            await login(mobile, otp);
            navigate('/');
        } catch (err) {
            setError('Invalid OTP. Please try again.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-slate-950 flex items-center justify-center p-4">
            <div className="max-w-md w-full bg-slate-900 rounded-2xl shadow-2xl border border-slate-800 p-8">
                <div className="text-center mb-8">
                    <h1 className="text-3xl font-bold text-white mb-2">Community Admin</h1>
                    <p className="text-slate-400">Manage your residents and services</p>
                </div>

                {error && (
                    <div className="bg-red-500/10 border border-red-500/50 text-red-500 p-3 rounded-lg mb-6 text-sm">
                        {error}
                    </div>
                )}

                {step === 1 ? (
                    <form onSubmit={handleSendOTP} className="space-y-4">
                        <div>
                            <label className="block text-slate-400 text-sm mb-2">Mobile Number</label>
                            <input
                                type="tel"
                                value={mobile}
                                onChange={(e) => setMobile(e.target.value)}
                                placeholder="9876543210"
                                className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
                                required
                            />
                        </div>
                        <button
                            type="submit"
                            disabled={loading}
                            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 rounded-lg transition-all disabled:opacity-50"
                        >
                            {loading ? 'Sending...' : 'Send OTP'}
                        </button>
                    </form>
                ) : (
                    <form onSubmit={handleVerifyOTP} className="space-y-4">
                        {receivedOtp && (
                            <div className="bg-blue-500/10 border border-blue-500/50 text-blue-400 p-3 rounded-lg mb-4 text-xs">
                                DEV MODE: Use OTP <strong>{receivedOtp}</strong> to login.
                            </div>
                        )}
                        <div>
                            <label className="block text-slate-400 text-sm mb-2">Verification Code</label>
                            <input
                                type="text"
                                value={otp}
                                onChange={(e) => setOtp(e.target.value)}
                                placeholder="123456"
                                className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all text-center tracking-widest"
                                required
                            />
                        </div>
                        <button
                            type="submit"
                            disabled={loading}
                            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 rounded-lg transition-all disabled:opacity-50"
                        >
                            {loading ? 'Verifying...' : 'Login'}
                        </button>
                        <button
                            type="button"
                            onClick={() => setStep(1)}
                            className="w-full text-slate-400 hover:text-white text-sm transition-all"
                        >
                            Change phone number
                        </button>
                    </form>
                )}
            </div>
        </div >
    );
};

export default Login;

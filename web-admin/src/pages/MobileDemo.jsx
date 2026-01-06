import React, { useState, useEffect } from 'react';
import PhoneFrame from '../components/PhoneFrame';
import {
    FaWallet,
    FaPlus,
    FaCamera,
    FaBell,
    FaStickyNote,
    FaFolder,
    FaCalendarAlt,
    FaHome,
    FaUsers,
    FaComments,
    FaUser,
    FaArrowLeft,
    FaUtensils,
    FaTaxi
} from 'react-icons/fa';

import { authService } from '../services/communityApi';

const MobileDemo = () => {
    const [screen, setScreen] = useState('splash');
    const [loading, setLoading] = useState(false);
    const [mobile, setMobile] = useState('');
    const [otp, setOtp] = useState('');
    const [receivedOtp, setReceivedOtp] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        const timer = setTimeout(() => setScreen('login'), 2000);
        return () => clearTimeout(timer);
    }, []);

    const handleSendOTP = async () => {
        if (!mobile) return setError('Enter mobile');
        setLoading(true);
        setError('');
        try {
            const data = await authService.sendOTP(mobile);
            if (data.otp) setReceivedOtp(data.otp);
            setScreen('otp-verify');
        } catch (err) {
            setError('Failed to send');
        } finally {
            setLoading(false);
        }
    };

    const handleVerifyOTP = async () => {
        setLoading(true);
        try {
            await authService.verifyOTP(mobile, otp);
            setScreen('dashboard');
        } catch (err) {
            setError('Invalid OTP');
        } finally {
            setLoading(false);
        }
    };

    const SplashScreen = () => (
        <div className="h-full flex flex-col items-center justify-center bg-slate-950">
            <div className="w-16 h-16 bg-blue-600 rounded-2xl flex items-center justify-center animate-bounce shadow-lg shadow-blue-500/20">
                <FaWallet className="text-3xl text-white" />
            </div>
            <h1 className="mt-4 text-white font-bold text-xl tracking-wider">SocWhiz</h1>
        </div>
    );

    const LoginScreen = () => (
        <div className="h-full px-6 flex flex-col justify-center bg-slate-950">
            <div className="mb-12 flex flex-col items-center">
                <FaWallet className="text-5xl text-blue-500 mb-4" />
                <h2 className="text-white text-2xl font-bold">Community Manager</h2>
                <p className="text-slate-400 text-xs mt-2 uppercase tracking-widest text-center">Personal Finance & Hub</p>
            </div>

            <div className="space-y-4">
                {error && <p className="text-red-500 text-[10px] text-center">{error}</p>}
                <div className="bg-slate-900 rounded-xl p-4 flex items-center gap-3 border border-white/5">
                    <FaUser className="text-slate-500" />
                    <input
                        type="text"
                        value={mobile}
                        onChange={(e) => setMobile(e.target.value)}
                        placeholder="Mobile Number"
                        className="bg-transparent border-none outline-none text-white text-sm w-full"
                        disabled={loading}
                    />
                </div>
                <button
                    onClick={handleSendOTP}
                    className="w-full bg-blue-600 hover:bg-blue-500 text-white font-bold py-4 rounded-xl shadow-lg shadow-blue-600/20 transition-all flex items-center justify-center gap-2"
                    disabled={loading}
                >
                    {loading ? 'Sending...' : 'Login with OTP'}
                </button>
            </div>
        </div>
    );

    const OTPVerifyScreen = () => (
        <div className="h-full px-6 flex flex-col justify-center bg-slate-950">
            <div className="mb-8 flex flex-col items-center">
                <h2 className="text-white text-xl font-bold">Verify OTP</h2>
                <p className="text-slate-400 text-[10px] mt-2 text-center">Sent to +91 {mobile}</p>
            </div>

            {receivedOtp && (
                <div className="bg-blue-500/10 border border-blue-500/50 text-blue-400 p-3 rounded-lg mb-4 text-[10px] text-center">
                    DEV MODE: Use OTP <strong>{receivedOtp}</strong>
                </div>
            )}

            <div className="space-y-4">
                <div className="bg-slate-900 rounded-xl p-4 border border-white/5">
                    <input
                        type="text"
                        value={otp}
                        onChange={(e) => setOtp(e.target.value)}
                        placeholder="000000"
                        className="bg-transparent border-none outline-none text-white text-xl w-full text-center tracking-[0.5em]"
                        maxLength={6}
                    />
                </div>
                <button
                    onClick={handleVerifyOTP}
                    className="w-full bg-blue-600 hover:bg-blue-500 text-white font-bold py-4 rounded-xl transition-all"
                    disabled={loading}
                >
                    {loading ? 'Verifying...' : 'Validate'}
                </button>
            </div>
        </div>
    );

    const ChatList = () => (
        <div className="h-full bg-slate-950 flex flex-col">
            <div className="p-4 flex items-center justify-between bg-slate-900/50 backdrop-blur-md sticky top-0 z-10 border-b border-white/5">
                <h2 className="text-white font-bold text-lg">Messages</h2>
                <FaPlus className="text-blue-500" />
            </div>
            <div className="flex-1 overflow-y-auto p-4 space-y-4 pb-24">
                {[
                    { name: 'Community Group', msg: 'The meeting will start at 8...', time: '10:45 AM', count: 5, color: 'bg-blue-500' },
                    { name: 'Alice Smith', msg: 'Sent you a document', time: 'Yesterday', count: 0, color: 'bg-indigo-500' },
                    { name: 'Bob Wilson', msg: 'Regarding the parking slot', time: '2 days ago', count: 0, color: 'bg-green-500' },
                ].map((chat, i) => (
                    <div key={i} className="flex items-center gap-4 bg-slate-900/50 p-3 rounded-2xl border border-white/5">
                        <div className={`w-12 h-12 rounded-2xl ${chat.color} flex items-center justify-center font-bold text-white`}>
                            {chat.name[0]}
                        </div>
                        <div className="flex-1 min-w-0">
                            <div className="flex justify-between items-center mb-1">
                                <h4 className="text-white text-sm font-bold truncate">{chat.name}</h4>
                                <span className="text-slate-500 text-[10px]">{chat.time}</span>
                            </div>
                            <p className="text-slate-400 text-xs truncate italic">"{chat.msg}"</p>
                        </div>
                        {chat.count > 0 && (
                            <div className="bg-blue-600 w-5 h-5 rounded-full flex items-center justify-center text-[10px] text-white font-bold">
                                {chat.count}
                            </div>
                        )}
                    </div>
                ))}
            </div>
            {/* Bottom Nav */}
            <div className="absolute bottom-0 left-0 right-0 p-3 bg-slate-900/80 backdrop-blur-lg border-t border-white/5 flex justify-around">
                <FaHome onClick={() => setScreen('dashboard')} className="text-slate-500 text-lg cursor-pointer" />
                <FaUsers className="text-slate-500 text-lg" />
                <FaComments className="text-blue-500 text-lg cursor-pointer" />
                <FaUser onClick={() => setScreen('profile')} className="text-slate-500 text-lg cursor-pointer" />
            </div>
        </div>
    );

    const CalendarScreen = () => (
        <div className="h-full bg-slate-950 flex flex-col">
            <div className="p-4 flex items-center gap-3 bg-slate-900/50 backdrop-blur-md sticky top-0 z-10 border-b border-white/5">
                <FaArrowLeft onClick={() => setScreen('dashboard')} className="text-white cursor-pointer" />
                <h2 className="text-white font-bold text-sm">Event Calendar</h2>
            </div>
            <div className="p-4 space-y-6 overflow-y-auto flex-1 pb-20">
                <div className="bg-slate-900 rounded-3xl p-6 border border-white/5">
                    <div className="flex justify-between items-center mb-4">
                        <span className="text-white font-bold text-xs">January 2026</span>
                    </div>
                    <div className="grid grid-cols-7 gap-2 text-center text-[8px] text-slate-500 font-bold uppercase mb-4">
                        {['Su', 'Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa'].map(d => <span key={d}>{d}</span>)}
                    </div>
                    <div className="grid grid-cols-7 gap-1 text-center">
                        {Array.from({ length: 31 }, (_, i) => (
                            <div key={i} className={`p-2 text-[10px] rounded-lg transition-all ${i + 1 === 5 ? 'bg-blue-600 text-white font-bold' : 'text-slate-400 hover:bg-slate-800'}`}>
                                {i + 1}
                            </div>
                        ))}
                    </div>
                </div>
                <div>
                    <h4 className="text-white text-xs font-bold mb-3">Upcoming Events</h4>
                    <div className="bg-blue-600/10 border border-blue-500/20 p-3 rounded-2xl flex items-center gap-3">
                        <div className="bg-blue-600 w-10 h-10 rounded-xl flex items-center justify-center text-white font-bold flex-shrink-0 text-xs text-center">
                            05 Jan
                        </div>
                        <div>
                            <p className="text-white text-xs font-bold">Annual General Meeting</p>
                            <p className="text-blue-400 text-[9px]">Today, 7:00 PM • Clubhouse</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );

    const ProfileScreen = () => (
        <div className="h-full bg-slate-950 flex flex-col">
            <div className="p-8 flex flex-col items-center bg-gradient-to-b from-blue-600/20 to-transparent">
                <div className="relative group">
                    <div className="w-20 h-20 rounded-[1.5rem] bg-blue-500 flex items-center justify-center text-2xl font-bold text-white border-4 border-slate-950 shadow-2xl">
                        JD
                    </div>
                    <div className="absolute -bottom-2 -right-2 w-7 h-7 bg-slate-900 rounded-xl flex items-center justify-center border-2 border-slate-950 text-blue-500">
                        <FaCamera className="text-[10px]" />
                    </div>
                </div>
                <h2 className="text-white text-base font-bold mt-4">John Doe</h2>
                <p className="text-slate-500 text-[10px]">Resident • Sunrise Gardens</p>
                <button className="mt-3 px-4 py-1.5 bg-slate-900 rounded-full text-[9px] font-bold text-blue-500 border border-blue-500/20">Edit Profile</button>
            </div>
            <div className="p-4 space-y-2 flex-1 overflow-y-auto pb-24">
                {[
                    { icon: FaWallet, label: 'Payment Methods', sub: 'VISA •••• 4242' },
                    { icon: FaUsers, label: 'My Community', sub: 'Sunrise Gardens Apts' },
                    { icon: FaBell, label: 'Notifications', sub: 'Sound, Vibrate, Alerts' },
                    { icon: FaUser, label: 'Security', sub: 'Password, Biometrics' },
                ].map((item, i) => (
                    <div key={i} className="flex items-center gap-4 bg-slate-900/30 p-3 rounded-2xl border border-white/5 active:bg-slate-800 transition-colors">
                        <div className="w-8 h-8 rounded-xl bg-slate-900 flex items-center justify-center text-slate-400 text-xs">
                            <item.icon />
                        </div>
                        <div className="flex-1 text-left">
                            <p className="text-white text-[11px] font-bold">{item.label}</p>
                            <p className="text-slate-500 text-[9px]">{item.sub}</p>
                        </div>
                    </div>
                ))}
                <button
                    onClick={() => setScreen('login')}
                    className="w-full mt-4 py-3 text-red-400 text-[10px] font-bold bg-red-500/5 rounded-2xl border border-red-500/10 active:bg-red-500/10"
                >
                    Logout Session
                </button>
            </div>
            {/* Bottom Nav */}
            <div className="absolute bottom-0 left-0 right-0 p-3 bg-slate-900/80 backdrop-blur-lg border-t border-white/5 flex justify-around">
                <FaHome onClick={() => setScreen('dashboard')} className="text-slate-500 text-lg cursor-pointer" />
                <FaUsers className="text-slate-500 text-lg" />
                <FaComments onClick={() => setScreen('chats')} className="text-slate-500 text-lg cursor-pointer" />
                <FaUser className="text-blue-500 text-lg cursor-pointer" />
            </div>
        </div>
    );

    const Dashboard = () => (
        <div className="h-full bg-slate-950 flex flex-col">
            <div className="p-4 flex justify-between items-center bg-slate-900/50 backdrop-blur-md sticky top-0 z-10">
                <div onClick={() => setScreen('profile')} className="flex items-center gap-2 cursor-pointer group">
                    <div className="w-8 h-8 rounded-full bg-blue-500 flex items-center justify-center text-[10px] font-bold text-white uppercase group-hover:ring-2 ring-blue-500/50 transition-all">JD</div>
                    <span className="text-white font-medium text-sm group-hover:text-blue-400 transition-colors">Hello, John!</span>
                </div>
                <FaBell className="text-slate-400 cursor-pointer hover:text-white transition-colors" />
            </div>

            <div className="flex-1 overflow-y-auto p-4 space-y-6 pb-20">
                {/* Balance Card */}
                <div className="bg-gradient-to-br from-blue-600 to-purple-600 p-5 rounded-2xl shadow-xl shadow-blue-500/10">
                    <p className="text-white/70 text-[10px] uppercase font-bold tracking-tighter">Total Balance</p>
                    <h3 className="text-white text-2xl font-bold mt-1">₹2,500.00</h3>
                    <div className="mt-4 flex gap-4">
                        <div>
                            <p className="text-white/60 text-[9px] uppercase">Income</p>
                            <p className="text-green-300 font-bold text-sm">₹3,600</p>
                        </div>
                        <div>
                            <p className="text-white/60 text-[9px] uppercase">Expenses</p>
                            <p className="text-red-300 font-bold text-sm">₹900</p>
                        </div>
                    </div>
                </div>

                {/* Quick Actions */}
                <div>
                    <h4 className="text-white text-sm font-bold mb-3">Quick Actions</h4>
                    <div className="grid grid-cols-3 gap-3">
                        {[
                            { icon: FaPlus, label: 'Add Expense', color: 'bg-orange-500/10 text-orange-500' },
                            { icon: FaCamera, label: 'Scan Bill', color: 'bg-blue-500/10 text-blue-500', action: () => setScreen('scanner') },
                            { icon: FaBell, label: 'Reminders', color: 'bg-purple-500/10 text-purple-500' },
                            { icon: FaStickyNote, label: 'Notes', color: 'bg-green-500/10 text-green-500' },
                            { icon: FaFolder, label: 'Documents', color: 'bg-indigo-500/10 text-indigo-500' },
                            { icon: FaCalendarAlt, label: 'Calendar', color: 'bg-red-500/10 text-red-500', action: () => setScreen('calendar') },
                        ].map((action, i) => (
                            <div
                                key={i}
                                onClick={action.action}
                                className="bg-slate-900 p-3 rounded-xl flex flex-col items-center justify-center gap-2 cursor-pointer hover:bg-slate-800 transition-all border border-white/5 active:scale-95"
                            >
                                <div className={`p-2 rounded-lg ${action.color}`}>
                                    <action.icon className="text-lg" />
                                </div>
                                <span className="text-slate-400 text-[9px] text-center font-medium leading-none">{action.label}</span>
                            </div>
                        ))}
                    </div>
                </div>

                {/* Transactions */}
                <div>
                    <div className="flex justify-between items-center mb-3">
                        <h4 className="text-white text-sm font-bold">Recent Transactions</h4>
                        <span className="text-blue-500 text-[10px] font-bold cursor-pointer">View All</span>
                    </div>
                    <div className="space-y-2">
                        {[
                            { icon: FaUtensils, label: 'Dinner Out', date: 'Today, 8:30 PM', amount: '-₹450', color: 'text-orange-400' },
                            { icon: FaTaxi, label: 'Uber Ride', date: 'Yesterday', amount: '-₹180', color: 'text-blue-400' },
                        ].map((tx, i) => (
                            <div key={i} className="bg-slate-900 p-3 rounded-xl flex items-center justify-between border border-white/5 text-left">
                                <div className="flex items-center gap-3">
                                    <div className={`w-9 h-9 rounded-lg bg-slate-800 flex items-center justify-center ${tx.color}`}>
                                        <tx.icon />
                                    </div>
                                    <div>
                                        <p className="text-white text-xs font-bold">{tx.label}</p>
                                        <p className="text-slate-500 text-[10px]">{tx.date}</p>
                                    </div>
                                </div>
                                <span className="text-red-400 font-bold text-xs">{tx.amount}</span>
                            </div>
                        ))}
                    </div>
                </div>
            </div>

            {/* Bottom Nav */}
            <div className="absolute bottom-0 left-0 right-0 p-3 bg-slate-900/80 backdrop-blur-lg border-t border-white/5 flex justify-around">
                <FaHome className="text-blue-500 text-lg cursor-pointer" />
                <FaUsers className="text-slate-500 text-lg cursor-pointer" />
                <FaComments onClick={() => setScreen('chats')} className="text-slate-500 text-lg cursor-pointer" />
                <FaUser onClick={() => setScreen('profile')} className="text-slate-500 text-lg cursor-pointer" />
            </div>
        </div>
    );

    const Scanner = () => (
        <div className="h-full bg-slate-950 flex flex-col relative overflow-hidden">
            <div className="p-4 flex items-center gap-3 bg-slate-900/50 backdrop-blur-md sticky top-0 z-10 text-white">
                <FaArrowLeft onClick={() => setScreen('dashboard')} className="cursor-pointer" />
                <span className="font-bold text-sm">Bill Scanner</span>
            </div>

            <div className="flex-1 flex flex-col items-center justify-center p-6 text-center space-y-6">
                <div className="relative w-full aspect-[3/4] border-2 border-dashed border-blue-500/50 rounded-2xl flex items-center justify-center overflow-hidden group">
                    <div className="absolute inset-4 border-2 border-blue-500 rounded-lg pointer-events-none after:absolute after:inset-0 after:bg-blue-500/5 after:animate-pulse"></div>
                    <div className="w-full h-1 bg-gradient-to-r from-transparent via-blue-400 to-transparent absolute top-1/2 left-0 animate-[scan_2s_infinite]"></div>
                    <p className="text-blue-400 text-[10px] font-bold uppercase tracking-widest z-10">Align bill within frame</p>
                </div>

                <div className="w-16 h-16 rounded-full bg-blue-600 flex items-center justify-center shadow-lg shadow-blue-600/20 active:scale-95 transition-all cursor-pointer">
                    <FaCamera className="text-white text-2xl" />
                </div>

                <div className="w-full bg-slate-900 p-4 rounded-2xl border border-white/5">
                    <p className="text-slate-500 text-[10px] uppercase font-bold text-left mb-2">Tips</p>
                    <ul className="text-slate-400 text-[9px] text-left space-y-1">
                        <li>• Ensure good lighting</li>
                        <li>• Flatten the bill before scanning</li>
                        <li>• Keep text clear and oriented properly</li>
                    </ul>
                </div>
            </div>
        </div>
    );

    return (
        <div className="min-h-screen bg-slate-950 flex flex-col items-center justify-center p-4">
            <div className="max-w-4xl w-full flex flex-col md:flex-row items-center gap-12">
                <div className="flex-1 text-center md:text-left space-y-4">
                    <div className="inline-block px-3 py-1 bg-blue-500/10 text-blue-500 rounded-full text-xs font-bold uppercase tracking-widest mb-2 border border-blue-500/20">Live Demo</div>
                    <h2 className="text-4xl font-extrabold text-white leading-tight">Interactive Mobile <br /><span className="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 to-purple-400 font-black">Experience</span></h2>
                    <p className="text-slate-400 max-w-sm">Explore the **SocWhiz Android App** right here in your browser. This simulator uses the exact UI components and theme logic implemented in our Flutter app.</p>
                    <div className="flex flex-wrap gap-3 justify-center md:justify-start pt-4">
                        <span className="px-3 py-1 bg-slate-900 rounded-lg text-xs border border-white/10 text-slate-300">✓ Dark Mode</span>
                        <span className="px-3 py-1 bg-slate-900 rounded-lg text-xs border border-white/10 text-slate-300">✓ OCR Simulation</span>
                        <span className="px-3 py-1 bg-slate-900 rounded-lg text-xs border border-white/10 text-slate-300">✓ Real-time Stats</span>
                    </div>
                </div>

                <div className="flex-1 relative group">
                    <div className="absolute -inset-10 bg-blue-600/5 blur-3xl rounded-full group-hover:bg-blue-600/10 transition-all duration-1000"></div>
                    <PhoneFrame>
                        {screen === 'splash' && <SplashScreen />}
                        {screen === 'login' && <LoginScreen />}
                        {screen === 'otp-verify' && <OTPVerifyScreen />}
                        {screen === 'dashboard' && <Dashboard />}
                        {screen === 'scanner' && <Scanner />}
                        {screen === 'chats' && <ChatList />}
                        {screen === 'calendar' && <CalendarScreen />}
                        {screen === 'profile' && <ProfileScreen />}
                    </PhoneFrame>
                </div>
            </div>

            <style>{`
        @keyframes scan {
          0%, 100% { transform: translateY(-100px); opacity: 0; }
          50% { opacity: 1; }
          100% { transform: translateY(100px); opacity: 0; }
        }
        .scrollbar-hide::-webkit-scrollbar { display: none; }
        .scrollbar-hide { -ms-overflow-style: none; scrollbar-width: none; }
      `}</style>
        </div>
    );
};

export default MobileDemo;

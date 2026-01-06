import React from 'react';
import { useAuth } from '../context/AuthContext';
import { FaBullhorn, FaPoll, FaHandshake, FaUsers } from 'react-icons/fa';

const Dashboard = () => {
    const { communityId } = useAuth();

    const stats = [
        { label: 'Total Members', value: '--', icon: <FaUsers />, color: 'bg-blue-600' },
        { label: 'Active Notices', value: '--', icon: <FaBullhorn />, color: 'bg-emerald-600' },
        { label: 'Active Polls', value: '--', icon: <FaPoll />, color: 'bg-purple-600' },
        { label: 'Recent Meetings', value: '--', icon: <FaHandshake />, color: 'bg-orange-600' },
    ];

    if (!communityId) {
        return (
            <div className="flex items-center justify-center h-full">
                <p className="text-slate-400">Please select a community from the header to begin.</p>
            </div>
        );
    }

    return (
        <div className="space-y-8">
            <div>
                <h1 className="text-3xl font-bold text-white mb-2">Welcome Back, Admin</h1>
                <p className="text-slate-400">Here's what's happening in your community today.</p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                {stats.map((stat) => (
                    <div key={stat.label} className="bg-slate-900 border border-slate-800 rounded-2xl p-6 shadow-lg">
                        <div className="flex items-center justify-between mb-4">
                            <div className={`${stat.color} p-3 rounded-xl text-white text-xl shadow-lg`}>
                                {stat.icon}
                            </div>
                        </div>
                        <div className="text-slate-400 text-sm">{stat.label}</div>
                        <div className="text-3xl font-bold text-white mt-1">{stat.value}</div>
                    </div>
                ))}
            </div>

            <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
                <div className="bg-slate-900 border border-slate-800 rounded-2xl p-6">
                    <h2 className="text-xl font-bold text-white mb-6">Quick Actions</h2>
                    <div className="grid grid-cols-2 gap-4">
                        <button className="bg-blue-600 hover:bg-blue-700 text-white p-4 rounded-xl flex flex-col items-center justify-center space-y-2 transition-all shadow-lg hover:shadow-blue-600/20">
                            <FaBullhorn className="text-2xl" />
                            <span className="font-medium">Post Notice</span>
                        </button>
                        <button className="bg-emerald-600 hover:bg-emerald-700 text-white p-4 rounded-xl flex flex-col items-center justify-center space-y-2 transition-all shadow-lg hover:shadow-emerald-600/20">
                            <FaPoll className="text-2xl" />
                            <span className="font-medium">Create Poll</span>
                        </button>
                        <button className="bg-purple-600 hover:bg-purple-700 text-white p-4 rounded-xl flex flex-col items-center justify-center space-y-2 transition-all shadow-lg hover:shadow-purple-600/20">
                            <FaHandshake className="text-2xl" />
                            <span className="font-medium">Add Meeting</span>
                        </button>
                        <button className="bg-slate-800 hover:bg-slate-750 text-white p-4 rounded-xl flex flex-col items-center justify-center space-y-2 transition-all shadow-lg">
                            <FaUsers className="text-2xl" />
                            <span className="font-medium">Manage Team</span>
                        </button>
                    </div>
                </div>

                <div className="bg-slate-900 border border-slate-800 rounded-2xl p-6">
                    <h2 className="text-xl font-bold text-white mb-6">Recent Activity</h2>
                    <div className="space-y-4">
                        <div className="text-slate-500 text-sm text-center py-8">
                            No recent activity to show for this community.
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Dashboard;

import React from 'react';
import { Outlet, Navigate } from 'react-router-dom';
import Sidebar from './Sidebar';
import { useAuth } from '../context/AuthContext';
import { communityService } from '../services/communityApi';

const Layout = () => {
    const { user, loading, communityId, selectCommunity } = useAuth();
    const [communities, setCommunities] = React.useState([]);

    React.useEffect(() => {
        if (user) {
            communityService.getCommunities().then((data) => {
                setCommunities(data.communities || []);
                if (!communityId && data.communities?.length > 0) {
                    selectCommunity(data.communities[0].id);
                }
            });
        }
    }, [user]);

    if (loading) return <div>Loading...</div>;
    if (!user) return <Navigate to="/login" />;

    return (
        <div className="flex h-screen bg-slate-950">
            <Sidebar />
            <div className="flex-1 flex flex-col overflow-hidden">
                <header className="bg-slate-900 h-16 border-b border-slate-800 flex items-center justify-between px-8 text-white">
                    <div className="flex items-center space-x-4">
                        <span className="text-slate-400">Managing:</span>
                        <select
                            value={communityId}
                            onChange={(e) => selectCommunity(e.target.value)}
                            className="bg-slate-800 border border-slate-700 rounded-md px-3 py-1 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                        >
                            <option value="">Select Community</option>
                            {communities.map((c) => (
                                <option key={c.id} value={c.id}>
                                    {c.name}
                                </option>
                            ))}
                        </select>
                    </div>
                    <div className="flex items-center space-x-4">
                        <span className="bg-blue-600/20 text-blue-400 px-3 py-1 rounded-full text-sm font-medium border border-blue-600/30">
                            Admin Mode
                        </span>
                    </div>
                </header>
                <main className="flex-1 overflow-x-hidden overflow-y-auto p-8">
                    <Outlet />
                </main>
            </div>
        </div>
    );
};

export default Layout;

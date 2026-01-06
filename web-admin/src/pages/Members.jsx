import React from 'react';

const Members = () => {
    return (
        <div className="space-y-6">
            <div>
                <h1 className="text-3xl font-bold text-white">Community Members</h1>
                <p className="text-slate-400">View and manage residents in your community</p>
            </div>

            <div className="bg-slate-900 border border-slate-800 rounded-2xl overflow-hidden shadow-lg">
                <table className="w-full text-left text-white border-collapse">
                    <thead>
                        <tr className="bg-slate-800 text-slate-400 text-sm uppercase tracking-wider">
                            <th className="p-4 border-b border-slate-700">Name</th>
                            <th className="p-4 border-b border-slate-700">Contact</th>
                            <th className="p-4 border-b border-slate-700">Role</th>
                            <th className="p-4 border-b border-slate-700 text-right">Joined</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr className="border-b border-slate-800 hover:bg-slate-800/30 transition-colors">
                            <td className="p-4 flex items-center space-x-3">
                                <div className="w-10 h-10 rounded-full bg-blue-600 flex items-center justify-center font-bold">A</div>
                                <span className="font-bold">Admin User</span>
                            </td>
                            <td className="p-4 text-slate-400">+91 9876543210</td>
                            <td className="p-4">
                                <span className="bg-blue-600/20 text-blue-400 px-3 py-1 rounded-full text-xs font-medium border border-blue-600/30">
                                    Admin
                                </span>
                            </td>
                            <td className="p-4 text-right text-slate-500">2026-01-01</td>
                        </tr>
                        <tr className="text-center">
                            <td colSpan="4" className="p-12 text-slate-500">
                                Member list integration coming soon.
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    );
};

export default Members;

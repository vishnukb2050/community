import React from 'react';
import { NavLink } from 'react-router-dom';
import { FaHome, FaBullhorn, FaPoll, FaHandshake, FaUsers, FaSignOutAlt } from 'react-icons/fa';
import { useAuth } from '../context/AuthContext';

const Sidebar = () => {
    const { logout } = useAuth();

    const menuItems = [
        { icon: <FaHome />, label: 'Dashboard', path: '/' },
        { icon: <FaBullhorn />, label: 'Notices', path: '/notices' },
        { icon: <FaPoll />, label: 'Polls', path: '/polls' },
        { icon: <FaHandshake />, label: 'Meetings', path: '/meetings' },
        { icon: <FaUsers />, label: 'Members', path: '/members' },
    ];

    return (
        <div className="w-64 bg-slate-900 h-screen text-white flex flex-col">
            <div className="p-6 text-2xl font-bold border-b border-slate-800">
                Admin Panel
            </div>
            <nav className="flex-1 p-4 space-y-2">
                {menuItems.map((item) => (
                    <NavLink
                        key={item.path}
                        to={item.path}
                        className={({ isActive }) =>
                            `flex items-center space-x-3 p-3 rounded-lg transition-colors ${isActive ? 'bg-blue-600 text-white' : 'hover:bg-slate-800 text-slate-400'
                            }`
                        }
                    >
                        {item.icon}
                        <span>{item.label}</span>
                    </NavLink>
                ))}
            </nav>
            <button
                onClick={logout}
                className="m-4 flex items-center space-x-3 p-3 rounded-lg hover:bg-red-600 hover:text-white text-slate-400 transition-colors"
            >
                <FaSignOutAlt />
                <span>Logout</span>
            </button>
        </div>
    );
};

export default Sidebar;

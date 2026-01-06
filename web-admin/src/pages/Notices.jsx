import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { noticeService } from '../services/communityApi';
import { FaPlus, FaTrash, FaBullhorn } from 'react-icons/fa';

const Notices = () => {
    const { communityId } = useAuth();
    const [notices, setNotices] = useState([]);
    const [loading, setLoading] = useState(true);
    const [showModal, setShowModal] = useState(false);
    const [form, setForm] = useState({ title: '', content: '' });

    useEffect(() => {
        if (communityId) loadNotices();
    }, [communityId]);

    const loadNotices = async () => {
        setLoading(true);
        try {
            const data = await noticeService.getNotices(communityId);
            setNotices(data.notices || []);
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    const handleCreate = async (e) => {
        e.preventDefault();
        try {
            await noticeService.createNotice({ ...form, community_id: communityId });
            setShowModal(false);
            setForm({ title: '', content: '' });
            loadNotices();
        } catch (err) {
            alert('Failed to create notice');
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Are you sure you want to delete this notice?')) {
            try {
                await noticeService.deleteNotice(id);
                loadNotices();
            } catch (err) {
                alert('Failed to delete notice');
            }
        }
    };

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold text-white">Notice Board</h1>
                    <p className="text-slate-400">Manage announcements for your community</p>
                </div>
                <button
                    onClick={() => setShowModal(true)}
                    className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-xl font-bold flex items-center space-x-2 shadow-lg hover:shadow-blue-600/20 transition-all"
                >
                    <FaPlus />
                    <span>Post Notice</span>
                </button>
            </div>

            {loading ? (
                <div className="text-white">Loading...</div>
            ) : (
                <div className="grid gap-4">
                    {notices.length === 0 ? (
                        <div className="bg-slate-900 border border-slate-800 rounded-2xl p-12 text-center">
                            <FaBullhorn className="text-slate-700 text-5xl mx-auto mb-4" />
                            <p className="text-slate-500">No notices posted yet.</p>
                        </div>
                    ) : (
                        notices.map((notice) => (
                            <div key={notice.id} className="bg-slate-900 border border-slate-800 rounded-2xl p-6 hover:border-slate-700 transition-all shadow-lg">
                                <div className="flex justify-between items-start">
                                    <div className="flex-1">
                                        <h3 className="text-xl font-bold text-white mb-2">{notice.title}</h3>
                                        <p className="text-slate-400 whitespace-pre-wrap">{notice.content}</p>
                                        <div className="text-slate-500 text-sm mt-4">
                                            Posted on: {new Date(notice.created_at).toLocaleDateString()}
                                        </div>
                                    </div>
                                    <button
                                        onClick={() => handleDelete(notice.id)}
                                        className="p-3 text-slate-500 hover:text-red-500 transition-colors"
                                    >
                                        <FaTrash />
                                    </button>
                                </div>
                            </div>
                        ))
                    )}
                </div>
            )}

            {showModal && (
                <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
                    <div className="bg-slate-900 border border-slate-800 rounded-2xl max-w-xl w-full p-8 shadow-2xl">
                        <h2 className="text-2xl font-bold text-white mb-6">Create New Notice</h2>
                        <form onSubmit={handleCreate} className="space-y-4">
                            <div>
                                <label className="block text-slate-400 text-sm mb-2">Notice Title</label>
                                <input
                                    type="text"
                                    value={form.title}
                                    onChange={(e) => setForm({ ...form, title: e.target.value })}
                                    className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                                    required
                                />
                            </div>
                            <div>
                                <label className="block text-slate-400 text-sm mb-2">Content</label>
                                <textarea
                                    value={form.content}
                                    onChange={(e) => setForm({ ...form, content: e.target.value })}
                                    rows="6"
                                    className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                                    required
                                ></textarea>
                            </div>
                            <div className="flex space-x-4 pt-4">
                                <button
                                    type="submit"
                                    className="flex-1 bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 rounded-lg shadow-lg"
                                >
                                    Publish Notice
                                </button>
                                <button
                                    type="button"
                                    onClick={() => setShowModal(false)}
                                    className="flex-1 bg-slate-800 hover:bg-slate-750 text-white font-bold py-3 rounded-lg"
                                >
                                    Cancel
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Notices;

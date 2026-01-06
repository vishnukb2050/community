import React, { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { pollService } from '../services/communityApi';
import { FaPlus, FaPoll, FaChevronRight, FaCheckCircle } from 'react-icons/fa';

const Polls = () => {
    const { communityId } = useAuth();
    const [polls, setPolls] = useState([]);
    const [loading, setLoading] = useState(true);
    const [showModal, setShowModal] = useState(false);
    const [form, setForm] = useState({ question: '', options: ['', ''] });

    useEffect(() => {
        if (communityId) loadPolls();
    }, [communityId]);

    const loadPolls = async () => {
        setLoading(true);
        try {
            const data = await pollService.getPoll(communityId);
            setPolls(data.polls || []);
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    const handleCreate = async (e) => {
        e.preventDefault();
        try {
            await pollService.createPoll({ ...form, community_id: communityId });
            setShowModal(false);
            setForm({ question: '', options: ['', ''] });
            loadPolls();
        } catch (err) {
            alert('Failed to create poll');
        }
    };

    const addOption = () => {
        setForm({ ...form, options: [...form.options, ''] });
    };

    const removeOption = (index) => {
        if (form.options.length > 2) {
            const newOptions = form.options.filter((_, i) => i !== index);
            setForm({ ...form, options: newOptions });
        }
    };

    return (
        <div className="space-y-6">
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold text-white">Polls & Voting</h1>
                    <p className="text-slate-400">Gather resident feedback on important matters</p>
                </div>
                <button
                    onClick={() => setShowModal(true)}
                    className="bg-emerald-600 hover:bg-emerald-700 text-white px-6 py-3 rounded-xl font-bold flex items-center space-x-2 shadow-lg hover:shadow-emerald-600/20 transition-all"
                >
                    <FaPlus />
                    <span>Create Poll</span>
                </button>
            </div>

            {loading ? (
                <div className="text-white">Loading...</div>
            ) : (
                <div className="grid gap-6">
                    {polls.length === 0 ? (
                        <div className="bg-slate-900 border border-slate-800 rounded-2xl p-12 text-center">
                            <FaPoll className="text-slate-700 text-5xl mx-auto mb-4" />
                            <p className="text-slate-500">No polls active at the moment.</p>
                        </div>
                    ) : (
                        polls.map((poll) => (
                            <div key={poll.id} className="bg-slate-900 border border-slate-800 rounded-2xl p-6 shadow-lg">
                                <h3 className="text-xl font-bold text-white mb-4">{poll.question}</h3>
                                <div className="space-y-3">
                                    {/* In production, fetch options for each poll */}
                                    <div className="text-slate-500 italic text-sm">Select poll to view results...</div>
                                </div>
                                <div className="flex justify-between items-center mt-6 pt-4 border-t border-slate-800 text-slate-500 text-sm">
                                    <span>Created: {new Date(poll.created_at).toLocaleDateString()}</span>
                                    <span className="flex items-center text-emerald-500">
                                        <FaCheckCircle className="mr-2" /> Active
                                    </span>
                                </div>
                            </div>
                        ))
                    )}
                </div>
            )}

            {showModal && (
                <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
                    <div className="bg-slate-900 border border-slate-800 rounded-2xl max-w-xl w-full p-8 shadow-2xl">
                        <h2 className="text-2xl font-bold text-white mb-6">Create New Poll</h2>
                        <form onSubmit={handleCreate} className="space-y-4">
                            <div>
                                <label className="block text-slate-400 text-sm mb-2">Question</label>
                                <input
                                    type="text"
                                    value={form.question}
                                    onChange={(e) => setForm({ ...form, question: e.target.value })}
                                    className="w-full bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500"
                                    placeholder="e.g. Should we paint the clubhouse blue?"
                                    required
                                />
                            </div>
                            <div className="space-y-3">
                                <label className="block text-slate-400 text-sm">Options</label>
                                {form.options.map((opt, idx) => (
                                    <div key={idx} className="flex space-x-2">
                                        <input
                                            type="text"
                                            value={opt}
                                            onChange={(e) => {
                                                const newOpts = [...form.options];
                                                newOpts[idx] = e.target.value;
                                                setForm({ ...form, options: newOpts });
                                            }}
                                            className="flex-1 bg-slate-800 border border-slate-700 rounded-lg px-4 py-3 text-white focus:outline-none focus:ring-2 focus:ring-emerald-500"
                                            placeholder={`Option ${idx + 1}`}
                                            required
                                        />
                                        {form.options.length > 2 && (
                                            <button
                                                type="button"
                                                onClick={() => removeOption(idx)}
                                                className="text-slate-500 hover:text-red-500"
                                            >
                                                <FaTrash />
                                            </button>
                                        )}
                                    </div>
                                ))}
                                <button
                                    type="button"
                                    onClick={addOption}
                                    className="text-emerald-500 hover:text-emerald-400 text-sm font-medium flex items-center"
                                >
                                    <FaPlus className="mr-2" /> Add another option
                                </button>
                            </div>
                            <div className="flex space-x-4 pt-4">
                                <button
                                    type="submit"
                                    className="flex-1 bg-emerald-600 hover:bg-emerald-700 text-white font-bold py-3 rounded-lg shadow-lg shadow-emerald-600/20"
                                >
                                    Launch Poll
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

export default Polls;

import { userNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { useState } from 'react';
import { useEffect } from 'react';

const Allbook = () => {
    // กำหนด State สำหรับจัดการข้อมูล
    const [Allbook, setAllbook] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchAllbook = async () => {
            try {
                setLoading(true);

                // เรียก API เพื่อดึงข้อมูลหนังสือทั้งหมด
                const response = await fetch('http://localhost:8080/api/v1/books');

                if (!response.ok) {
                    throw new Error('Failed to fetch all books');
                }

                const data = await response.json();
                setAllbook(data);
                setError(null);

            } catch (err) {
                setError(err.message);
                console.error('Error fetching all books:', err);

            } finally {
                setLoading(false);
            }
        };

        // เรียกใช้ฟังก์ชันดึงข้อมูล
        fetchAllbook();
    }, []); // [] = dependency array ว่าง = รันครั้งเดียว
    return (
        <div className="min-h-screen bg-sky-50">
            <nav className="bg-gradient-to-r from-sky-400 to-sky-500 text-white shadow-lg sticky top-0 z-50">
                <div className="container mx-auto px-6 py-4">
                    <div className="flex justify-between items-center">
                        <Link to="/" className="flex items-center space-x-3 group">
                            <div className="h-10 w-10 bg-white rounded-lg flex items-center justify-center group-hover:scale-110 transition-transform">
                                <span className="text-sky-500 font-bold text-xl">B</span>
                            </div>
                            <h1 className="text-2xl font-bold group-hover:text-sky-100 transition-colors">
                                BookStore Manager
                            </h1>
                        </Link>
                        <Link
                            to="/store-manager/add-book"
                            className="bg-white text-sky-600 px-6 py-2.5 rounded-lg font-semibold 
                                hover:bg-sky-50 hover:text-sky-700 hover:shadow-lg 
                                transform hover:scale-105 transition-all duration-200 
                                flex items-center space-x-2"
                        >
                            <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                            </svg>
                            <span>เพิ่มหนังสือใหม่</span>
                        </Link>
                    </div>
                </div>
            </nav>
            <div className="container mx-auto p-6">
                <div className="mb-8">
                    <h2 className="text-3xl font-bold text-gray-800 mb-2">จัดการหนังสือ</h2>
                    <p className="text-gray-600">จัดการข้อมูลหนังสือทั้งหมดในระบบ</p>
                </div>
                <div className="bg-white rounded-xl shadow-lg overflow-hidden">
                    <div className="overflow-x-auto">
                        <table className="min-w-full">
                            <thead className="bg-gradient-to-r from-sky-400 to-sky-500 text-white">
                                <tr>
                                    <th className="py-3 px-6 text-left">ID</th>
                                    <th className="py-3 px-6 text-left">ชื่อหนังสือ</th>
                                    <th className="py-3 px-6 text-left">ผู้แต่ง</th>
                                    <th className="py-3 px-6 text-left">หมวดหมู่</th>
                                    <th className="py-3 px-6 text-left">ราคา</th>
                                    <th className="py-3 px-6 text-left">จำนวนหน้า</th>
                                    <th className="py-3 px-6 text-left">จัดการ</th>
                                </tr>
                            </thead>
                            <tbody>
                                {loading ? (
                                    <tr>
                                        <td colSpan="7" className="py-8 text-center text-gray-500">
                                            กำลังโหลดข้อมูล...
                                        </td>
                                    </tr>
                                ) : error ? (
                                    <tr>
                                        <td colSpan="7" className="py-8 text-center text-red-500">
                                            เกิดข้อผิดพลาด: {error}
                                        </td>
                                    </tr>
                                ) : Allbook.length === 0 ? (
                                    <tr>
                                        <td colSpan="7" className="py-8 text-center text-gray-500">
                                            ไม่พบข้อมูลหนังสือ
                                        </td>
                                    </tr>
                                ) : (
                                    Allbook.map((book) => (
                                        <tr key={book.id} className="border-b hover:bg-sky-50 transition-colors">
                                            <td className="py-3 px-6">{book.id}</td>
                                            <td className="py-3 px-6 font-medium">{book.title}</td>
                                            <td className="py-3 px-6">{book.author}</td>
                                            <td className="py-3 px-6">
                                                <span className="inline-block px-3 py-1 bg-sky-100 text-sky-700 rounded-full text-sm">
                                                    {book.category}
                                                </span>
                                            </td>
                                            <td className="py-3 px-6 font-semibold text-sky-600">
                                                ฿{book.price.toFixed(2)}
                                            </td>
                                            <td className="py-3 px-6">{book.pages || 'N/A'}</td>
                                            <td className="py-3 px-6">
                                                <div className="flex space-x-2">
                                                    <button className="bg-sky-500 text-white px-4 py-2 rounded-lg hover:bg-sky-600 transition-colors duration-200">
                                                        แก้ไข
                                                    </button>
                                                    <button className="bg-rose-500 text-white px-4 py-2 rounded-lg hover:bg-rose-600 transition-colors duration-200">
                                                        ลบ
                                                    </button>
                                                </div>
                                            </td>
                                        </tr>
                                    ))
                                )}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Allbook;
import { userNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { useState } from 'react';
import { useEffect } from 'react';


function Allbook() {
    return (
        <div className=" ">
            <navbar className="p-4 text-black shadow-lg">
                <div className="container mx-auto flex justify-between items-center">
                    <h1 className="text-2xl font-bold pl-4">Bookstore</h1>
                    <Link to="/store-manager/add-book" className="bg-white text-viridian-600 px-4 py-2 rounded-lg hover:bg-gray-100 transition-colors duration-200">
                        เพิ่มหนังสือใหม่
                    </Link>
                </div>
            </navbar>
            <div className="container mx-auto p-4">
                <h2 className="text-3xl font-bold mb-6 p-4">จัดการหนังสือ</h2>
                <div className="overflow-x-auto">
                    <table className="min-w-full bg-white shadow-md rounded-lg overflow-hidden">
                        <thead className="bg-viridian-600 text-black">
                            <tr>
                                <th className="py-3 px-6 text-left">รหัสหนังสือ</th>
                                <th className="py-3 px-6 text-left">ชื่อหนังสือ</th>
                                <th className="py-3 px-6 text-left">ผู้แต่ง</th>
                                <th className="py-3 px-6 text-left">หมวดหมู่</th>
                                <th className="py-3 px-6 text-left">ราคา</th>
                                <th className="py-3 px-6 text-left">จำนวน</th>
                                <th className="py-3 px-6 text-left">จัดการ</th>
                            </tr>
                        </thead>
                        <tbody>
                            {/* Sample Data Row */}
                            <tr className="border-b hover:bg-gray-100">
                                <td className="py-3 px-6">B001</td>
                                <td className="py-3 px-6">The Great Gatsby</td>
                                <td className="py-3 px-6">F. Scott Fitzgerald</td>
                                <td className="py-3 px-6">Fiction</td>
                                <td className="py-3 px-6">300</td>
                                <td className="py-3 px-6">50</td>
                                <td className="py-3 px-6">
                                    <button className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-600 transition-colors duration-200 mr-2">
                                        แก้ไข
                                    </button>
                                    <button className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition-colors duration-200">
                                        ลบ
                                    </button>
                                </td>
                            </tr>
                            {/* Add more rows as needed */}
                        </tbody>
                    </table>
                </div>
            </div>
        </div >

    );
}

export default Allbook;
import React, { useState, useEffect } from 'react';
import { Pagination } from 'antd'; // 导入 Pagination 组件
import './SearchComponent.css';

// 假设我们有以下省份和部门的数据结构
const provincesData = [
    {id: 0, name: 'All'},
    {id: 35, name: '中央'},
    {id: 1, name: '北京市'},
    {id: 2, name: '天津市'},
    {id: 3, name: '河北省'},
    {id: 4, name: '山西省'},
    {id: 5, name: '内蒙古自治区'},
    {id: 6, name: '辽宁省'},
    {id: 7, name: '吉林省'},
    {id: 8, name: '黑龙江省'},
    {id: 9, name: '上海市'},
    {id: 10, name: '江苏省'},
    {id: 11, name: '浙江省'},
    {id: 12, name: '安徽省'},
    {id: 13, name: '福建省'},
    {id: 14, name: '江西省'},
    {id: 15, name: '山东省'},
    {id: 16, name: '河南省'},
    {id: 17, name: '湖北省'},
    {id: 18, name: '湖南省'},
    {id: 19, name: '广东省'},
    {id: 20, name: '广西壮族自治区'},
    {id: 21, name: '海南省'},
    {id: 22, name: '重庆市'},
    {id: 23, name: '四川省'},
    {id: 24, name: '贵州省'},
    {id: 25, name: '云南省'},
    {id: 26, name: '西藏自治区'},
    {id: 27, name: '陕西省'},
    {id: 28, name: '甘肃省'},
    {id: 29, name: '青海省'},
    {id: 30, name: '宁夏回族自治区'},
    {id: 31, name: '新疆维吾尔自治区'},
    {id: 32, name: '台湾省'},
    {id: 33, name: '香港特别行政区'},
    {id: 34, name: '澳门特别行政区'}
];

const departmentsData = [
    {id: 0, name: 'All'},
    {id: 1, name: '科学技术部'},
    {id: 2, name: '教育部'},
    {id: 3, name: '工信部'},
];

// 生成选项列表的函数
const generateOptions = (dataList) => dataList.map(item => (
    <option key={item.id} value={item.id}>{item.name}</option>
));

// 根据ID获取省份或部门名称的函数
const getNameById = (dataList, id) => dataList.find(item => item.id === id)?.name;

const SearchComponent = () => {
    const [query, setQuery] = useState('');
    const [pid, setPid] = useState(0); // 初始 pid 值为 0
    const [did, setDid] = useState(0); // 初始 did 值为 0
    const [results, setResults] = useState([]);
    const [currentPage, setCurrentPage] = useState(1); // 将当前页数初始值设为 1
    const [total, setTotal] = useState(0); // 总条目数

    useEffect(() => {
        fetchData();
    }, [currentPage]); // 当当前页数、pid 或 did 发生变化时，重新获取数据

    const handleChange = (event) => {
        setQuery(event.target.value);
    };

    const handlePidChange = (event) => {
        setPid(event.target.value);
    };

    const handleDidChange = (event) => {
        setDid(event.target.value);
    };

    const formatDate = (dateString) => {
        const date = new Date(dateString);
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const day = String(date.getDate()).padStart(2, '0');
        return `${year}-${month}-${day}`;
    };

    const fetchData = async () => {
        try {
            const response = await fetch(`http://localhost:3000/search?s=${query}&pid=${pid}&did=${did}&page=${currentPage}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setResults(data.hits);
            setTotal(data.total.value)
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        setCurrentPage(1); // 提交搜索时重置当前页数为第一页
        fetchData(); // 在重置 currentPage 后立即执行 fetchData
    };

    const handlePageChange = (page) => {
        setCurrentPage(page); // 当分页改变时更新当前页数
    };

    return (
        <div className="search-container">
            <form onSubmit={handleSubmit} className="search-form">
                <input
                    type="text"
                    placeholder="Search..."
                    value={query}
                    onChange={handleChange}
                    className="search-input"
                />
                <select value={pid} onChange={handlePidChange} className="search-select">
                    {generateOptions(provincesData)}
                </select>
                <select value={did} onChange={handleDidChange} className="search-select">
                    {generateOptions(departmentsData)}
                </select>
                <button type="submit" className="search-button">Search</button>
            </form>
            <ul className="search-results">
                {results.map((result, index) => (
                    <li key={index} className="search-result">
                        <a href={result._source.url} target="_blank" rel="noopener noreferrer"
                           className="result-title">{result._source.title}</a>
                        <p className="result-date">Date: {formatDate(result._source.date)}</p>
                        <p className="result-content" dangerouslySetInnerHTML={{__html: result.highlight.content}}/>
                        <p className="result-source-tag">
                            来源：<span className="result-province">{getNameById(provincesData, result._source.province_id)}</span> /
                            <span className="result-department">{getNameById(departmentsData, result._source.department_id)}</span>
                        </p>
                    </li>
                ))}
            </ul>
            <Pagination simple
                        onChange={handlePageChange}
                        current={currentPage}
                        total={total}
                        style={{marginTop: '20px', marginBottom: '20px', textAlign: 'center' }} />
        </div>
    );
};

export default SearchComponent;

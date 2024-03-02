import React, { useState, useEffect } from 'react';
import { Pagination } from 'antd'; // 导入 Pagination 组件
import './SearchComponent.css';

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
                    <option value={0}>All</option>
                    <option value={35}>中央</option>
                    <option value={1}>北京市</option>
                    <option value={2}>天津市</option>
                    <option value={3}>河北省</option>
                    <option value={4}>山西省</option>
                    <option value={5}>内蒙古自治区</option>
                    <option value={6}>辽宁省</option>
                    <option value={7}>吉林省</option>
                    <option value={8}>黑龙江省</option>
                    <option value={9}>上海市</option>
                    <option value={10}>江苏省</option>
                    <option value={11}>浙江省</option>
                    <option value={12}>安徽省</option>
                    <option value={13}>福建省</option>
                    <option value={14}>江西省</option>
                    <option value={15}>山东省</option>
                    <option value={16}>河南省</option>
                    <option value={17}>湖北省</option>
                    <option value={18}>湖南省</option>
                    <option value={19}>广东省</option>
                    <option value={20}>广西壮族自治区</option>
                    <option value={21}>海南省</option>
                    <option value={22}>重庆市</option>
                    <option value={23}>四川省</option>
                    <option value={24}>贵州省</option>
                    <option value={25}>云南省</option>
                    <option value={26}>西藏自治区</option>
                    <option value={27}>陕西省</option>
                    <option value={28}>甘肃省</option>
                    <option value={29}>青海省</option>
                    <option value={30}>宁夏回族自治区</option>
                    <option value={31}>新疆维吾尔自治区</option>
                    <option value={32}>台湾省</option>
                    <option value={33}>香港特别行政区</option>
                    <option value={34}>澳门特别行政区</option>
                </select>
                <select value={did} onChange={handleDidChange} className="search-select">
                    <option value={0}>All</option>
                    <option value={1}>科学技术部</option>
                    <option value={2}>教育部</option>
                    <option value={3}>工信部</option>
                </select>
                <button type="submit" className="search-button">Search</button>
            </form>
            <ul className="search-results">
                {results.map((result, index) => (
                    <li key={index} className="search-result">
                        <a href={result._source.url} target="_blank" rel="noopener noreferrer" className="result-title">{result._source.title}</a>
                        <p className="result-date">Date: {formatDate(result._source.date)}</p>
                        <p className="result-content" dangerouslySetInnerHTML={{__html: result.highlight.content}}/>
                    </li>
                ))}
            </ul>
            <Pagination simple
                        onChange={handlePageChange}
                        current={currentPage}
                        total={total}
                        style={{ marginTop: '20px', marginBottom: '20px', textAlign: 'center' }} />
        </div>
    );
};

export default SearchComponent;

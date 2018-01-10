#include <iostream>
#include <vector>
#include <chrono>

template<typename It, typename Compare>
auto mergeSort(It _first, It _end, Compare &_func)
{
	using T = std::iterator_traits<It>::value_type;
	auto distance = _end - _first;
	if (static_cast<int>(distance) <= 1)
	{
		return std::vector<T>(_first, _end);
	}
	else
	{
		auto middle = _end - static_cast<int>(distance / 2);

		auto Lprime = mergeSort(_first, middle, _func);
		auto Rprime = mergeSort(middle, _end, _func);

		auto merge = std::vector<T>();

		auto L_iter = Lprime.begin();
		auto R_iter = Rprime.begin();

		while (true)
		{
			if (L_iter == Lprime.end())
			{
				merge.insert(merge.end(), R_iter, Rprime.end());
				break;
			}
			if (R_iter == Rprime.end())
			{
				merge.insert(merge.end(), L_iter, Lprime.end());
				break;
			}

			if (_func(*L_iter, *R_iter) == true) //right precede left
			{
				merge.emplace_back(std::move(*R_iter));
				R_iter++;
			}
			else
			{
				merge.emplace_back(std::move(*L_iter));
				L_iter++;
			}
		}

		return merge;
	}
}

int main(void)
{

	auto compareFunc = [](int lhs, int rhs)->bool
	{
		if (lhs > rhs)
			return true;
		else
			return false;
	};

	std::vector<int> problemVector{
		4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2,
		5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3,
		6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4,
		7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5,
		8, 9, 10, 11, 12, 11, 10, 9, 8, 7, 6,
		7, 8, 9, 10, 11, 10, 9, 8, 7, 6, 5,
		6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4,
		5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3,
		4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2,
		3, 4, 5, 6, 7, 6, 5, 4, 3, 2, 1,
		2, 3, 4, 5, 6, 5, 4, 3, 2, 1, 0 };

	int index = 0;
	int square = static_cast<int>(std::sqrt(problemVector.size()));

	std::cout << "Before merge sort." << std::endl;
	for (auto& val : problemVector)
	{
		if (index % square == 0)
			std::cout << std::endl;

		std::cout << val << ", ";
		index++;
	}

	index = 0;

	auto start = std::chrono::steady_clock::now();

	auto result = mergeSort(problemVector.begin(), problemVector.end(), compareFunc);
	std::cout << "\nAfter merge sort." << std::endl;
	for (auto& val : result)
	{
		if (index % square == 0)
			std::cout << std::endl;

		std::cout << val << ", ";
		index++;
	}

	auto end = std::chrono::steady_clock::now();

	auto diff = end - start;
	std::cout << "\nMerge sort 걸린 시간 : " << std::chrono::duration <double, std::milli>(diff).count() << " (ms)" << std::endl;

	system("pause");
	return 0;
}
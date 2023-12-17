
// fifth bruteforce solution, time C. O(n), space C. O(n)
type FoodRatings struct {

	food_dict   map[string]SubContent
	hcuis_dict   map[string]SubContent
	cuisine_dict map[string]map[string]int
}

type SubContent struct {
	value string
	rating   int
}

func (this *FoodRatings)add_hcuisine(food, cuisine string, rating int){
	 if (this.hcuis_dict[cuisine].rating < rating){
		this.hcuis_dict[cuisine] = SubContent{
			food,
			rating,
		}
	} else if ((this.hcuis_dict[cuisine].rating > rating) && (this.hcuis_dict[cuisine].value == food)){
		lmax := 0
		for k, v :=range this.cuisine_dict[cuisine]{
			if v > lmax {
			this.hcuis_dict[cuisine] = SubContent{
				k,
				v,
			}
			lmax = v
			}else if v == lmax {
				if this.hcuis_dict[cuisine].value < k {
					this.hcuis_dict[cuisine] = SubContent{
					this.hcuis_dict[cuisine].value,
					v,
					}
				}else {
					this.hcuis_dict[cuisine] = SubContent{
						k,
						v,
					}
				}

			}
		}
	} 
	 if (this.hcuis_dict[cuisine].rating == rating){
		if this.hcuis_dict[cuisine].value < food {
			this.hcuis_dict[cuisine] = SubContent{
				this.hcuis_dict[cuisine].value,
				rating,
			}
		}else {
		this.hcuis_dict[cuisine] = SubContent{
			food,
			rating,
		}
	}
	}

}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {

	food_dict := make(map[string]SubContent)
	cuisine_dict := make(map[string]map[string]int)
	hcuis_dict := make(map[string]SubContent)

	for i,v := range foods {
		food_dict[v] = SubContent{
			cuisines[i],
			ratings[i],
		}
		if _, ok :=cuisine_dict[cuisines[i]]; !ok{
			cuisine_dict[cuisines[i]]= map[string]int{}
			cuisine_dict[cuisines[i]][v]=ratings[i]
		}else {
		cuisine_dict[cuisines[i]][v]=ratings[i]
		}

		if _, ok := hcuis_dict[cuisines[i]]; !ok {
		hcuis_dict[cuisines[i]] = SubContent{
			v,
			ratings[i],
		}
		} else if (hcuis_dict[cuisines[i]].rating < ratings[i]){
		hcuis_dict[cuisines[i]] = SubContent{
			v,
			ratings[i],
		}
		} else if (hcuis_dict[cuisines[i]].rating == ratings[i]){
			if hcuis_dict[cuisines[i]].value < foods[i] {
			hcuis_dict[cuisines[i]] = SubContent{
				hcuis_dict[cuisines[i]].value,
				ratings[i],
			}
		}else {
		hcuis_dict[cuisines[i]] = SubContent{
			foods[i],
			ratings[i],
		}
	}
		}
	}

	return  FoodRatings{
		food_dict: food_dict,
		hcuis_dict: hcuis_dict,
		cuisine_dict: cuisine_dict,
	}

}

func (this *FoodRatings) ChangeRating(food string, newRating int)  {
	entry := this.food_dict[food]
	entry.rating = newRating
	this.cuisine_dict[entry.value][food] = newRating
	this.add_hcuisine(food, this.food_dict[food].value, newRating)   
}


func (this *FoodRatings) HighestRated(cuisine string) string {
	fmt.Println(this.hcuis_dict)
	return this.hcuis_dict[cuisine].value
}


/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
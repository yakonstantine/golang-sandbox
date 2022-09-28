package srv

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golangbase/tasks/hangman_online/srv/core/ent"
	"golangbase/tasks/hangman_online/srv/core/mocks"
	"testing"
	"time"
)

func TestWordService_GetWord_GetMoreThenBuffer(t *testing.T) {
	// arrange
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wCnt := 7
	buffSize := wCnt - 2
	wgm := new(mocks.WordGetterMock)
	for i := 0; i < buffSize*2; i++ {
		expW := ent.Word(fmt.Sprintf("foo%d", i+1))
		wgm.On("GetWord").Return(expW, nil).Once()
	}

	sut := NewWordService(ctx, wgm, buffSize)

	// act
	for i := 0; i < wCnt; i++ {
		expW := ent.Word(fmt.Sprintf("foo%d", i+1))
		w, err := sut.GetWord()
		// assert
		if !assert.Nil(t, err) || !assert.Equal(t, w, expW) {
			return
		}
	}
}

func TestWordService_GetWord_GetLessThenBuffer(t *testing.T) {
	// arrange
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	wCnt := 3
	buffSize := 5
	wgm := new(mocks.WordGetterMock)
	for i := 0; i < buffSize+wCnt; i++ {
		expW := ent.Word(fmt.Sprintf("foo%d", i+1))
		wgm.On("GetWord").Return(expW, nil).Once()
	}

	sut := NewWordService(ctx, wgm, buffSize)

	// act
	for i := 0; i < wCnt; i++ {
		expW := ent.Word(fmt.Sprintf("foo%d", i+1))
		w, err := sut.GetWord()
		// assert
		if !assert.Nil(t, err) || !assert.Equal(t, w, expW) {
			return
		}
	}
}

func TestWordService_GetWord_ShouldGetErrorFromClosedChan(t *testing.T) {
	// arrange
	ctx, cancel := context.WithCancel(context.TODO())

	buffSize := 5
	wgm := new(mocks.WordGetterMock)
	for i := 0; i < buffSize*2; i++ {
		expW := ent.Word(fmt.Sprintf("foo%d", i+1))
		wgm.On("GetWord").Return(expW, nil).Once()
	}

	sut := NewWordService(ctx, wgm, buffSize)
	// act
	cancel()
	time.Sleep(100 * time.Millisecond)
	//assert
	for i := 0; i < buffSize*2; i++ {
		w, err := sut.GetWord()
		t.Log(w)
		if err != nil {
			return
		}
	}
	assert.Fail(t, "The channel didn't close")
}
